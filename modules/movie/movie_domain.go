package users

import (
	"test-xsis/modules/movie/model"
	"test-xsis/schemas"
)

type Repository interface {
	CreateMovieRepository(movie *model.Movie) (error, int64)
	GetMoviesRepository() (*[]model.Movie, error)
	GetMovieByIDRepository(id int64) (*model.Movie, error)
	GetMovieByTitleRepository(title string) (*model.Movie, error)
	UpdateMovieByIDRepository(id int64, movie *model.Movie) error
	DeleteMovieByIDRepository(id int64) error
}

type Service interface {
	CreateMovieService(movie schemas.CreateMovieRequest) (error, int64)
	GetMoviesService() (*[]model.Movie, error)
	GetMovieByIDService(id int64) (*model.Movie, error)
	GetMovieByTitleService(title string) (*model.Movie, error)
	UpdateMovieByIDService(id int64, movie schemas.UpdateMovieRequest) error
	DeleteMovieByIDService(id int64) error
}
