package godeez

import (
	"encoding/json"
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

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tracks struct {
		Data []Track `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tracks); err != nil {
		return nil, err
	}

	return tracks.Data, nil
}