package handlers

import (
	"fmt"
	"movie_recommender/services"
)

func StartCLI() {
	for {
		fmt.Println("\nMovie Recommender")
		fmt.Println("1. Get Movie Recommendations")
		fmt.Println("2. List Watch Later")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Show available genres
			fmt.Println("Available Genres:")
			for id, name := range services.GenreMap {
				fmt.Printf("ID: %s, Name: %s\n", id, name)
			}

			fmt.Print("Enter a genre ID or name: ")
			var input string
			fmt.Scanln(&input)

			var genreID string
			var found bool

			// Check if input is a genre ID or name
			if len(input) > 2 {
				genreID, found = services.GetGenreID(input)
			} else {
				genreID, found = input, true
			}

			if !found {
				fmt.Println("Invalid genre. Please try again.")
				continue
			}

			// Fetch movies based on genre
			movies, err := services.FetchMovies(genreID)
			if err != nil {
				fmt.Println("Error fetching movies:", err)
			} else {
				fmt.Println("Recommended Movies:")
				for i, movie := range movies {
					fmt.Printf("ID: %d | Title: %v\n", i+1, movie["title"])
				}

				// Prompt user to add a movie to "Watch Later"
				fmt.Print("\nEnter the ID of the movie you want to add to 'Watch Later' list (or press 0 to skip): ")
				var movieChoice int
				fmt.Scanln(&movieChoice)

				if movieChoice > 0 && movieChoice <= len(movies) {
					selectedMovie := movies[movieChoice-1]
					title := selectedMovie["title"].(string)
					movieID := int(selectedMovie["id"].(float64))
					category := "Movie"

					err := services.AddToWatchLater(1, movieID, title, category)
					if err != nil {
						fmt.Println("Error adding to 'Watch Later' list:", err)
					} else {
						fmt.Println("Movie added to 'Watch Later' list successfully!")
					}
				} else if movieChoice == 0 {
					fmt.Println("No movie added to 'Watch Later'.")
				} else {
					fmt.Println("Invalid choice. Please try again.")
				}
			}

		case 2:
			// List movies in "Watch Later"
			fmt.Println("Your 'Watch Later' List:")
			watchLaterList, err := services.ListWatchLater(1)
			if err != nil {
				fmt.Println("Error retrieving 'Watch Later' list:", err)
			} else {
				for _, movie := range watchLaterList {
					fmt.Printf("ID: %v | Title: %v | Category: %v\n",
						movie["movie_id"], movie["title"], movie["category"])
				}
			}

		case 3:
			// Exit the program
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
