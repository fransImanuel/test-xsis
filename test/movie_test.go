package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"test-xsis/modules/movie/handler"
	"test-xsis/modules/movie/model"
	"test-xsis/schemas"
	"test-xsis/test/mocks"
	"test-xsis/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMovieHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a mock movie service
	mockMovieService := new(mocks.MovieServiceMock)

	// Set up the gin engine and handler
	r := gin.Default()
	h := &handler.MovieHandler{
		MovieService: mockMovieService,
	}

	// Create mock response for the movie service
	mockMovieService.On("CreateMovieService", mock.Anything).Return(nil, int64(1))

	// Create a request body
	movieReq := schemas.CreateMovieRequest{
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Rating:      8.8,
		Image:       "inception.jpg",
	}
	body, _ := json.Marshal(movieReq)

	// Create a new request
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/movie/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	r.POST("/api/v1/movie/create", h.CreateMovieHandler)
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	mockMovieService.AssertExpectations(t)
}

func TestGetAllMoviesHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//Create a mock movie service
	mockMovieService := new(mocks.MovieServiceMock)

	//Set up the gin engine and handler
	r := gin.Default()
	h := &handler.MovieHandler{
		MovieService: mockMovieService,
	}

	//Create mock data for the movies
	movies := []model.Movie{
		{
			FullAudit: schemas.FullAudit{
				ID: 1,
			},
			Title:       utils.StrPtr("Inception"),
			Description: utils.StrPtr("A mind-bending thriller"),
			Rating:      utils.Float64Ptr(8.8),
			Image:       utils.StrPtr("inception.jpg"),
		},
		{
			FullAudit: schemas.FullAudit{
				ID: 1,
			},
			Title:       utils.StrPtr("The Dark Knight"),
			Description: utils.StrPtr("A superhero crime drama"),
			Rating:      utils.Float64Ptr(9.0),
			Image:       utils.StrPtr("darkknight.jpg"),
		},
	}

	// Mock the GetMoviesService method
	mockMovieService.On("GetMoviesService").Return(&movies, nil)

	// Create a new request
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/movies", nil)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	//Serve the HTTP request
	r.GET("/api/v1/movies", h.GetAllMoviesHandler)
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	mockMovieService.AssertExpectations(t)
}

func TestGetMovieByIDHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockMovieService := new(mocks.MovieServiceMock)

	r := gin.Default()
	h := &handler.MovieHandler{
		MovieService: mockMovieService,
	}

	//Create mock data for the movies
	movies := &model.Movie{
		FullAudit: schemas.FullAudit{
			ID: 1,
		},
		Title:       utils.StrPtr("Inception"),
		Description: utils.StrPtr("A mind-bending thriller"),
		Rating:      utils.Float64Ptr(8.8),
		Image:       utils.StrPtr("inception.jpg"),
	}

	mockMovieService.On("GetMovieByIDService", int64(1)).Return(movies, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/movie/1", nil)

	w := httptest.NewRecorder()

	r.GET("/api/v1/movie/:id", h.GetMovieByIDHandler)
	r.ServeHTTP(w, req)

	// Parse the response body
	var responseMovie schemas.Response
	_ = json.Unmarshal(w.Body.Bytes(), &responseMovie)

	extractMovie := responseMovie.Data.(map[string]interface{})
	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Inception", extractMovie["title"])
	assert.Equal(t, "A mind-bending thriller", extractMovie["description"])
	assert.Equal(t, 8.8, extractMovie["rating"])
	assert.Equal(t, "inception.jpg", extractMovie["image"])

	// Verify that the mock's expectations were met
	mockMovieService.AssertExpectations(t)

}

func TestUpdateMovieByIDHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockMovieService := new(mocks.MovieServiceMock)

	r := gin.Default()
	h := &handler.MovieHandler{
		MovieService: mockMovieService,
	}

	// Create mock data for the update request
	updateReq := schemas.UpdateMovieRequest{
		Title:       "Inception Updated",
		Description: "An updated mind-bending thriller",
		Rating:      9.0,
		Image:       "inception_updated.jpg",
	}

	// Mock the UpdateMovieByIDService call
	mockMovieService.On("UpdateMovieByIDService", int64(1), updateReq).Return(nil)

	reqBody, _ := json.Marshal(updateReq)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/movie/1", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.PUT("/api/v1/movie/:id", h.UpdateMovieByIDHandler)
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	// Optionally check response message if needed

	// Verify that the mock's expectations were met
	mockMovieService.AssertExpectations(t)
}

func TestDeleteMovieByIDHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockMovieService := new(mocks.MovieServiceMock)

	r := gin.Default()
	h := &handler.MovieHandler{
		MovieService: mockMovieService,
	}

	// Mock the DeleteMovieByIDService call
	mockMovieService.On("DeleteMovieByIDService", int64(1)).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/movie/1", nil)

	w := httptest.NewRecorder()

	r.DELETE("/api/v1/movie/:id", h.DeleteMovieByIDHandler)
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify that the mock's expectations were met
	mockMovieService.AssertExpectations(t)
}
