package repository

import (
	"github.com/viniblima/atfilms/models"
	"gorm.io/gorm"
)

type uploadRepository struct {
	Db *gorm.DB
}

type UploadRepository interface {
}

func (r *uploadRepository) CreatePhoto(c *models.Photo) (*models.Photo, error) {
	err := r.Db.Create(c).Error
	return c, err
}

func (r *uploadRepository) CreateVideo(c *models.Video) (*models.Video, error) {
	err := r.Db.Create(c).Error
	return c, err
}

func (repo *uploadRepository) RemovePhotoByID(id string) error {
	var photo models.Photo
	err := repo.Db.Model(&models.Photo{}).Where("ID = ?", id).Delete(photo).Error
	return err
}

func (repo *uploadRepository) RemoveVideoByID(id string) error {
	var video models.Video
	err := repo.Db.Model(&models.Video{}).Where("ID = ?", id).Delete(video).Error
	return err
}
