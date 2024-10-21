package handler

import (
	"fmt"
	"net/http"
	"strconv"
	movie "test-xsis/modules/movie"
	"test-xsis/schemas"
	"test-xsis/utils"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	MovieService movie.Service
}

func InitMovieHandler(g *gin.Engine, itemService movie.Service) {
	handler := &MovieHandler{
		MovieService: itemService,
	}

	routeAPI := g.Group("/Movie")
	routeAPI.GET("", handler.GetAllMoviesHandler)
	routeAPI.GET("/:id", handler.GetMovieByIDHandler)
	routeAPI.GET("/title/:title", handler.GetMovieByTitleHandler)
	routeAPI.POST("", handler.CreateMovieHandler)
	routeAPI.PATCH("/:id", handler.UpdateMovieByIDHandler)
	routeAPI.DELETE("/:id", handler.DeleteMovieByIDHandler)
}

// Create Movie
// @Tags Movies
// @Summary Create or Add New Movie
// @Description Create Movie
// @ID Movie-Create
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.CreateMovieRequest true "body data"
// @Success 200  {object} schemas.Response
// @Router /Movie [post]
func (h *MovieHandler) CreateMovieHandler(c *gin.Context) {
	var req schemas.CreateMovieRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Required field is empty", nil)
		return
	}

	err, ID := h.MovieService.CreateMovieService(req)
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Create Movie", map[string]interface{}{
		"id": ID,
	})

}

// Get Movie
// @Tags Movies
// @Summary Get List of Movie
// @Description Get Movie
// @ID Movie-Get
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200  {object} schemas.Response
// @Router /Movie [get]
func (h *MovieHandler) GetAllMoviesHandler(c *gin.Context) {
	movie, err := h.MovieService.GetMoviesService()
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Get Movies", movie)
}

// Get Movie By ID
// @Tags Movies
// @Summary Get Movie By ID
// @Description Get Movie By id
// @ID MovieByID-Get
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "Movie ID"
// @Success 200  {object} schemas.Response
// @Router /Movie/{id} [get]
func (h *MovieHandler) GetMovieByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Param Path Empty", nil)
		return
	}

	intValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", fmt.Sprintf("Error converting string to int64: %v", err), nil)
		return
	}

	movie, err := h.MovieService.GetMovieByIDService(intValue)
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Get Movie By ID", movie)
}

// Get Movie By Title
// @Tags Movies
// @Summary Get Movie By Title
// @Description Get Movie By id
// @Title MovieByTitle-Get
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param        id   path      string  true  "Movie Title"
// @Success 200  {object} schemas.Response
// @Router /Movie/{title} [get]
func (h *MovieHandler) GetMovieByTitleHandler(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Param Path Empty", nil)
		return
	}

	movie, err := h.MovieService.GetMovieByTitleService(title)
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Get Movie By ID", movie)
}

// Update Movie By ID
// @Tags Movies
// @Summary Update Movie By ID
// @Description Update Movie By id
// @ID MovieByID-Update
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "Movie ID"
// @Param data body schemas.UpdateMovieRequest true "body data"
// @Success 200  {object} schemas.Response
// @Router /Movie/{id} [patch]
func (h *MovieHandler) UpdateMovieByIDHandler(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Param Path Empty", nil)
		return
	}

	var req schemas.UpdateMovieRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Required field is empty", nil)
		return
	}

	intValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", fmt.Sprintf("Error converting string to int64: %v", err), nil)
		return
	}

	if err := h.MovieService.UpdateMovieByIDService(intValue, req); err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}

	utils.APIResponse(c, http.StatusOK, "success", "Success Update Movie By ID", nil)
}

// Delete Movie By ID
// @Tags Movies
// @Summary Delete Movie By ID
// @Description Delete Movie By id
// @ID MovieByID-Delete
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "Movie ID"
// @Success 200  {object} schemas.Response
// @Router /Movie/{id} [delete]
func (h *MovieHandler) DeleteMovieByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Param Path Empty", nil)
		return
	}

	intValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", fmt.Sprintf("Error converting string to int64: %v", err), nil)
		return
	}

	if err := h.MovieService.DeleteMovieByIDService(intValue); err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Delete Movie By ID", nil)
}
