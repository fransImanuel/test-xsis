// movie_service_mock.go
package mocks

import (
	"test-xsis/modules/movie/model"
	"test-xsis/schemas"

	"github.com/stretchr/testify/mock"
)

// MovieServiceMock is a mock implementation of the MovieService interface.
type MovieServiceMock struct {
	mock.Mock
}

func (m *MovieServiceMock) CreateMovieService(req schemas.CreateMovieRequest) (error, int64) {
	args := m.Called(req)
	return args.Error(0), args.Get(1).(int64)
}

func (m *MovieServiceMock) GetMoviesService() (*[]model.Movie, error) {
	args := m.Called()
	return args.Get(0).(*[]model.Movie), args.Error(1)
}

func (m *MovieServiceMock) GetMovieByIDService(id int64) (*model.Movie, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Movie), args.Error(1)
}

func (m *MovieServiceMock) UpdateMovieByIDService(id int64, req schemas.UpdateMovieRequest) error {
	args := m.Called(id, req)
	return args.Error(0)
}

func (m *MovieServiceMock) DeleteMovieByIDService(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}
