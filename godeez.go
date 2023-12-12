package godeez

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const deezerAPIBaseURL = "https://api.deezer.com"

// Track represents the structure of a Deezer track.
type Track struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist struct {
		Name string `json:"name"`
	} `json:"artist"`
	// Add other fields as needed
}

// GetTopTracks retrieves the top tracks from the Deezer charts.
func GetTopTracks() ([]Track, error) {
	url := deezerAPIBaseURL + "/chart/tracks"

	fmt.Println("Fetching track data from API")
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var tracks struct {
		Data []Track `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tracks); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return tracks.Data, nil
}