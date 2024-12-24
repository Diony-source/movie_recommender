package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const tmdbBaseURL = "https://api.themoviedb.org/3"

// FetchMovies fetches movies from TMDb API based on genre ID
func FetchMovies(genreID string) ([]map[string]interface{}, error) {
	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("TMDB_API_KEY is not set in environment variables")
	}

	url := fmt.Sprintf("%s/discover/movie?api_key=%s&with_genres=%s", tmdbBaseURL, apiKey, genreID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch movies: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	movies, ok := result["results"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var movieList []map[string]interface{}
	for _, movie := range movies {
		movieList = append(movieList, movie.(map[string]interface{}))
	}

	return movieList, nil
}
