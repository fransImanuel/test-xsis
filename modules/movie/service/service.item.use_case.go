package service

import (
	movies "test-xsis/modules/movie"
	"test-xsis/modules/movie/model"
	"test-xsis/schemas"
)

type MovieService struct {
	MoviesRepository movies.Repository
}

func InitMoviesService(MoviesRepository movies.Repository) movies.Service {
	return &MovieService{
		MoviesRepository: MoviesRepository,
	}
}

func (i *MovieService) CreateMovieService(movie schemas.CreateMovieRequest) (error, int64) {
	itemModel := &model.Movie{
		Title:       &movie.Title,
		Description: &movie.Description,
		Rating:      &movie.Rating,
		Image:       &movie.Image,
	}
	err, id := i.MoviesRepository.CreateMovieRepository(itemModel)
	if err != nil {
		return err, 0
	}

	return nil, id
}

func (i *MovieService) GetMoviesService() (*[]model.Movie, error) {
	movie, err := i.MoviesRepository.GetMoviesRepository()
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (i *MovieService) GetMovieByIDService(id int64) (*model.Movie, error) {
	item, err := i.MoviesRepository.GetMovieByIDRepository(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *MovieService) GetMovieByTitleService(title string) (*model.Movie, error) {
	item, err := i.MoviesRepository.GetMovieByTitleRepository(title)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *MovieService) UpdateMovieByIDService(id int64, movie schemas.UpdateMovieRequest) error {
	itemModel := &model.Movie{
		Title:       &movie.Title,
		Description: &movie.Description,
		Rating:      &movie.Rating,
		Image:       &movie.Image,
	}
	if err := i.MoviesRepository.UpdateMovieByIDRepository(id, itemModel); err != nil {
		return err
	}

	return nil
}

func (i *MovieService) DeleteMovieByIDService(id int64) error {
	if err := i.MoviesRepository.DeleteMovieByIDRepository(id); err != nil {
		return err
	}

	return nil
}
