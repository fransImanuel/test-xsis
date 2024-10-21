package repository

import (
	"errors"
	"test-xsis/constant"
	movies "test-xsis/modules/movie"
	"test-xsis/modules/movie/model"
	"test-xsis/schemas"

	"gorm.io/gorm"
)

type MovieRepository struct {
	DBPostgres *gorm.DB
	//DBMongoDB
	//DBMinio, etc
}

func InitMovieRepository(db *gorm.DB) movies.Repository {
	return &MovieRepository{
		DBPostgres: db,
	}
}

func (u *MovieRepository) CreateMovieRepository(movie *model.Movie) (error, int64) {
	db := u.DBPostgres

	movie.InitAudit(constant.OPERATION_SQL_INSERT)

	results := db.Create(&movie)
	if results.Error != nil {
		return results.Error, 0
	}

	return nil, movie.ID

}

func (u *MovieRepository) GetMoviesRepository() (*[]model.Movie, error) {
	var movies *[]model.Movie
	db := u.DBPostgres

	// Get all records
	results := db.Find(&movies)
	// SELECT * FROM ITEMS;
	if results.Error != nil {
		return nil, results.Error
	}

	return movies, nil
}

func (u *MovieRepository) GetMovieByIDRepository(id int64) (*model.Movie, error) {
	var movies model.Movie
	db := u.DBPostgres

	// Get all records
	results := db.Unscoped().First(&movies, id)
	// SELECT * FROM ITEMS;
	if results.Error != nil {
		return nil, results.Error
	}

	return &movies, nil
}

func (u *MovieRepository) GetMovieByTitleRepository(title string) (*model.Movie, error) {
	var movies model.Movie
	db := u.DBPostgres

	// Get all records
	results := db.Unscoped().First(&movies, `"Title" = ? `, title)
	// SELECT * FROM ITEMS;
	if results.Error != nil {
		return nil, results.Error
	}

	return &movies, nil
}

func (u *MovieRepository) UpdateMovieByIDRepository(id int64, movie *model.Movie) error {
	var movies model.Movie
	db := u.DBPostgres

	movie.InitAudit(constant.OPERATION_SQL_UPDATE)

	result := db.Model(&movies).Where("id = ?", id).Updates(model.Movie{Title: movie.Title, Description: movie.Description, Rating: movie.Rating, Image: movie.Image, FullAudit: schemas.FullAudit{
		ModifiedTime: movie.ModifiedTime,
	}})

	if result.RowsAffected < 1 {
		return errors.New("No Data Affected")
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *MovieRepository) DeleteMovieByIDRepository(id int64) error {
	db := u.DBPostgres

	if err := db.Delete(&model.Movie{}, id).Error; err != nil {
		return err
	}

	return nil
}
