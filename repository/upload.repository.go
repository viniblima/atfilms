package repository

import (
	"github.com/viniblima/atfilms/database"
	"github.com/viniblima/atfilms/models"
	"gorm.io/gorm"
)

type uploadRepository struct {
	Db *gorm.DB
}

type UploadRepository interface {
	CreatePhoto(c *models.Photo) (*models.Photo, error)
	CreateVideo(c *models.Video) (*models.Video, error)
	RemovePhoto(c *models.Photo) error
	RemoveVideo(c *models.Video) error
}

func (r *uploadRepository) CreatePhoto(c *models.Photo) (*models.Photo, error) {
	err := r.Db.Create(c).Error
	return c, err
}

func (r *uploadRepository) CreateVideo(c *models.Video) (*models.Video, error) {
	err := r.Db.Create(c).Error
	return c, err
}

func (repo *uploadRepository) RemovePhoto(photo *models.Photo) error {
	err := repo.Db.Delete(photo).Error
	return err
}

func (repo *uploadRepository) RemoveVideo(video *models.Video) error {
	err := repo.Db.Delete(video).Error
	return err
}

func NewUploadRepository() UploadRepository {
	return &uploadRepository{
		Db: database.Db,
	}
}
