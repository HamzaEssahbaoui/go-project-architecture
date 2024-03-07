package handlers

import (
	"best-architecture/internal/models"
	"best-architecture/internal/services"
	"html/template"
	"log"
	"net/http"
)

// HomeHandler struct will hold any dependencies needed by the home page, such as the movie service.
type HomeHandler struct {
	Template     *template.Template
	MovieService services.MovieDBService
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Movies      []models.Movie
		SearchQuery string
		Title       string
		Header      string
	}
	// Example values
	data.Title = "Movie Finder"
	data.Header = "Find Your Favorite Movies"
	// Check if there's a search query.

	data.SearchQuery = r.URL.Query().Get("keyword")
	if data.SearchQuery != "" {
		log.Println("Search Query:", data.SearchQuery) // Log only when a query is present

		// Fetch movies based on the search query.
		movies, err := h.MovieService.SearchMovies(data.SearchQuery)
		if err != nil {
			log.Printf("Error searching for movies: %v", err)
			http.Error(w, "Failed to search for movies", http.StatusInternalServerError)
			return
		}

		log.Println("Found Movies:", len(movies)) // Debugging: log the number of movies found
		data.Movies = movies
	}

	// Execute the template with the search results.
	if err := h.Template.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

// NewHomeHandler initializes a new HomeHandler with the necessary dependencies.
func NewHomeHandler(movieService services.MovieDBService) *HomeHandler {
	// Combine base template with the specific page template
	tmpl, err := template.ParseFiles("web/templates/base.gohtml", "web/templates/home.gohtml")
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}

	return &HomeHandler{
		Template:     tmpl,
		MovieService: movieService,
	}
}
