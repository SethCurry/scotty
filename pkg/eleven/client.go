package eleven

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Reference: https://elevenlabs.io/docs/api-reference/getting-started

type Client struct {
	apiKey string
}

type ListVoicesResponse struct {
	Voices []VoicesListItem `json:"voices"`
}

type VoicesListItem struct {
	Name        string `json:"name"`
	VoiceID     string `json:"voice_id"`
	Category    string `json:"category"`
	Description string `json:"description"`
	PreviewURL  string `json:"preview_url"`
}

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

func (c *Client) makeJSONRequest(method string, url string, body interface{}) (*http.Request, error) {
	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %w", err)
	}

	return c.makeRequest(method, url, reqBody)
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}

func (c *Client) withJSONResponse(req *http.Request, output interface{}) error {
	resp, err := c.doRequest(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(output)
	return err
}

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
	Stability       int  `json:"stability"`
	SimilarityBoost int  `json:"similarity_boost"`
	Style           int  `json:"style"`
	UseSpeakerBoost bool `json:"use_speaker_boost"`
}

type TTSOptions struct {
	VoiceSettings VoiceSettings
}

type TTSRequest struct {
	Text          string        `json:"text"`
	VoiceSettings VoiceSettings `json:"voice_settings"`
}

func (c *Client) TTS(text string, voiceID string, output io.Writer, settings VoiceSettings) error {
	reqBody := &TTSRequest{
		Text:          text,
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

	_, err = io.Copy(output, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to copy data: %w", err)
	}

	return nil
}
