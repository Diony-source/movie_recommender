package services

// GenreMap contains TMDb genre IDs and names
var GenreMap = map[string]string{
	"28":    "Action",
	"12":    "Adventure",
	"16":    "Animation",
	"35":    "Comedy",
	"80":    "Crime",
	"99":    "Documentary",
	"18":    "Drama",
	"10751": "Family",
	"14":    "Fantasy",
	"36":    "History",
	"27":    "Horror",
	"10402": "Music",
	"9648":  "Mystery",
	"10749": "Romance",
	"878":   "Science Fiction",
	"10770": "TV Movie",
	"53":    "Thriller",
	"10752": "War",
	"37":    "Western",
}

// GetGenreName returns the genre name for a given ID
func GetGenreName(genreID string) (string, bool) {
	name, exists := GenreMap[genreID]
	return name, exists
}

// GetGenreID returns the genre ID for a given name
func GetGenreID(genreName string) (string, bool) {
	for id, name := range GenreMap {
		if name == genreName {
			return id, true
		}
	}
	return "", false
}
