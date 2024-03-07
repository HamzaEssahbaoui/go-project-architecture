package services

import (
	"best-architecture/internal/models"
	"encoding/json"
	"fmt"

	"net/http"
	"net/url"
)

const (
	baseURL        = "https://api.themoviedb.org/3"
	searchEndpoint = "/search/movie"
	movieEndpoint  = "/movie/"
)

type MovieDBService interface {
	SearchMovies(keyword string) ([]models.Movie, error)
	FetchMovieDetails(movieID string) (*models.Movie, error)
}

type movieDBClient struct {
	apiKey     string
	httpClient *http.Client
}

func NewMovieDBService(apiKey string, httpClient *http.Client) MovieDBService {
	return &movieDBClient{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (c *movieDBClient) SearchMovies(keyword string) ([]models.Movie, error) {
	requestURL := fmt.Sprintf("%s%s?api_key=%s&query=%s", baseURL, searchEndpoint, c.apiKey, url.QueryEscape(keyword))
	resp, err := c.httpClient.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Results []models.Movie `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c *movieDBClient) FetchMovieDetails(movieID string) (*models.Movie, error) {
	requestURL := fmt.Sprintf("%s%s%s?api_key=%s", baseURL, movieEndpoint, movieID, c.apiKey)
	resp, err := c.httpClient.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var detail models.Movie
	if err := json.NewDecoder(resp.Body).Decode(&detail); err != nil {
		return nil, err
	}

	return &detail, nil
}
