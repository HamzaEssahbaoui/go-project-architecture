package handlers

import (
	"best-architecture/internal/models"
	"best-architecture/internal/services"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// MovieHandler struct will contain dependencies for the movie detail page, such as the movie service.
type MovieHandler struct {
	Template     *template.Template
	MovieService services.MovieDBService
}

func NewMovieHandler(movieService services.MovieDBService) *MovieHandler {
	// Combine base template with the specific page template
	tmpl, err := template.ParseFiles("web/templates/base.gohtml", "web/templates/movie.gohtml")
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}

	return &MovieHandler{
		Template:     tmpl,
		MovieService: movieService,
	}
}

func (m *MovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}
	movieID := pathParts[2]

	completeMovie, err := m.MovieService.FetchMovieDetails(movieID)
	if err != nil {
		log.Printf("Failed to fetch movie details: %v", err)
		http.Error(w, "Failed to fetch movie details", http.StatusInternalServerError)
		return
	}

	// Adjust here: Wrap movieDetail into a structure that matches the template expectation.
	data := struct {
		Movie  *models.Movie // Ensure this type matches what FetchMovieDetails returns
		Title  string
		Header string
	}{
		Movie: completeMovie,
		Title: completeMovie.Title,
		// Assuming you want the movie's title as the page title
		Header: "Movie Details", // Set a generic header or use movieDetail.Title
	}

	if err := m.Template.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
