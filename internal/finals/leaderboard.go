package finals

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// LeaderboardResponse is the response from a leaderboard request to
// The Finals API.
type LeaderboardResponse struct {
	Data []LeaderboardPlayer `json:"data"`
}

// LeaderboardPlayer is a single player in the leaderboard.
type LeaderboardPlayer struct {
	Name   string `json:"name"`
	Rank   int    `json:"rank"`
	League string `json:"league"`
}

// Leaderboard queries The Finals API for leaderboard results.
func Leaderboard(name string) (*LeaderboardPlayer, error) {
	// https://api.the-finals-leaderboard.com/v1/leaderboard/s3/crossplay?name=anarchy

	u, err := url.Parse("https://api.the-finals-leaderboard.com/v1/leaderboard/s3/crossplay")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("name", name)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	var respData LeaderboardResponse
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}

	if len(respData.Data) == 0 {
		return nil, errors.New("no player found")
	} else if len(respData.Data) > 1 {
		return nil, errors.New("multiple players found with same name")
	}

	return &respData.Data[0], nil
}
