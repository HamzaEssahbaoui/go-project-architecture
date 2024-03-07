package models

// Movie represents the basic information about a movie.
type Movie struct {
	Overview    string `json:"overview"`
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
}
