package eleven

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Reference: https://elevenlabs.io/docs/api-reference/getting-started

// ClientOption is a function that configures a Client instance.
type ClientOption func(*Client)

// WithHTTPClient configures the HTTP client used by the Client instance.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// NewClient creates a new *Client instance.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	client := &Client{
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}

	for _, v := range opts {
		v(client)
	}

	return client
}

// Client is a client for the Eleven API.  It stores
// common configuration options like the API key to use
// and the HTTP client to use for requests.
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// ListVoicesResponse stores the response from a call to ListVoices.
type ListVoicesResponse struct {
	Voices []VoicesListItem `json:"voices"`
}

// VoicesListItem is a single voice from the ListVoices call.
type VoicesListItem struct {
	Name        string `json:"name"`
	VoiceID     string `json:"voice_id"`
	Category    string `json:"category"`
	Description string `json:"description"`
	PreviewURL  string `json:"preview_url"`
}

// makeRequest creates a new HTTP request and sets the Authorization, Content-Type
// and Accept headers for the request.  This provides a stable source of pre-configured
// HTTP requests.
func (c *Client) makeRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("Accept", "application/json")

	return req, nil
}

// makeJSONRequest does the same as makeRequest, but encodes a provided struct as JSON
// to use in the request's body.
func (c *Client) makeJSONRequest(method string, url string, body interface{}) (*http.Request, error) {
	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %w", err)
	}

	return c.makeRequest(method, url, reqBody)
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

// withJSONResponse makes a request to the provided URL and decodes its response body as JSON into
// the provided output interface{}.
func (c *Client) withJSONResponse(req *http.Request, output interface{}) error {
	resp, err := c.doRequest(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(output)
	return err
}

// ListVoices calls the Eleven API to get a list of voices for the current account.
func (c *Client) ListVoices() (*ListVoicesResponse, error) {
	req, err := c.makeRequest("GET", "https://api.elevenlabs.com/v1/voices", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var resp ListVoicesResponse

	err = c.withJSONResponse(req, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	return &resp, nil
}

type VoiceSettings struct {
	Stability       float32 `json:"stability"`
	SimilarityBoost float32 `json:"similarity_boost"`
	Style           int     `json:"style"`
	UseSpeakerBoost bool    `json:"use_speaker_boost"`
}

type TTSOptions struct {
	VoiceSettings VoiceSettings
}

type TTSRequest struct {
	Text          string        `json:"text"`
	ModelID       string        `json:"model_id"`
	VoiceSettings VoiceSettings `json:"voice_settings"`
}

// TTS calls the Eleven API to synthesize speech for a given text and voice ID.
// The output will be written to the provided io.Writer.
//
// This call is synchronous, and can potentially take a long time to complete.
func (c *Client) TTS(text string, voiceID string, output io.Writer, settings VoiceSettings) error {
	reqBody := &TTSRequest{
		Text:          text,
		ModelID:       "eleven_multilingual_v2",
		VoiceSettings: VoiceSettings{},
	}

	req, err := c.makeJSONRequest("POST", fmt.Sprintf("https://api.elevenlabs.io/v1/text-to-speech/%s", voiceID), reqBody)
	if err != nil {
		return err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		return fmt.Errorf("invalid status code: %d: %s", resp.StatusCode, string(body))
	}

	_, err = io.Copy(output, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to copy data: %w", err)
	}

	return nil
}
