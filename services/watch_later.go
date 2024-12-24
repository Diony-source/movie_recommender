package services

import (
	"context"
	"movie_recommender/database"
)

// AddToWatchLater adds a movie to the "watch later" list
func AddToWatchLater(userID, movieID int, title, category string) error {
	query := "INSERT INTO watch_later (user_id, movie_id, title, category) VALUES ($1, $2, $3, $4)"
	_, err := database.DB.Exec(context.Background(), query, userID, movieID, title, category)
	return err
}

// ListWatchLater retrieves all movies in the "watch later" list for a user
func ListWatchLater(userID int) ([]map[string]interface{}, error) {
	query := "SELECT movie_id, title, category FROM watch_later WHERE user_id = $1"
	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var watchLaterList []map[string]interface{}
	for rows.Next() {
		var movieID int
		var title, category string
		err := rows.Scan(&movieID, &title, &category)
		if err != nil {
			return nil, err
		}
		watchLaterList = append(watchLaterList, map[string]interface{}{
			"movie_id": movieID,
			"title":    title,
			"category": category,
		})
	}

	return watchLaterList, nil
}
