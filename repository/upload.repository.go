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
	UpdatePhoto(c *models.Photo) (*models.Photo, error)
	CreateVideo(c *models.Video) (*models.Video, error)
	UpdateVideo(c *models.Video) (*models.Video, error)
	RemovePhotoByID(id string) error
	RemoveVideoByID(id string) error
	GetPhotosBySliderID(id string) (*[]models.Photo, error)
	GetPhotoByID(id string) (*models.Photo, error)
	GetVideoByID(id string) (*models.Video, error)
	GetVideosByComponentID(id string) (*[]models.Video, error)
}

func (r *uploadRepository) CreatePhoto(c *models.Photo) (*models.Photo, error) {
	err := r.Db.Create(c).Error
	return c, err
}

func (r *uploadRepository) UpdatePhoto(c *models.Photo) (*models.Photo, error) {
	err := r.Db.Save(c).Error
	return c, err
}

func (r *uploadRepository) CreateVideo(c *models.Video) (*models.Video, error) {
	err := r.Db.Create(c).Error
	return c, err
}

func (r *uploadRepository) UpdateVideo(c *models.Video) (*models.Video, error) {
	err := r.Db.Save(c).Error
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

func (repo *uploadRepository) GetPhotosBySliderID(id string) (*[]models.Photo, error) {
	var ps []models.Photo
	err := repo.Db.Where("slider_id = ?", id).Find(&ps).Error
	return &ps, err
}

func (repo *uploadRepository) GetVideosByComponentID(id string) (*[]models.Video, error) {
	var ls []models.Video
	err := repo.Db.Where("job_component_videos_id = ?", id).Find(&ls).Error
	return &ls, err
}

func (repo *uploadRepository) GetPhotoByID(id string) (*models.Photo, error) {
	var p models.Photo
	err := repo.Db.Where("ID = ?", id).First(&p).Error
	return &p, err
}

func (repo *uploadRepository) GetVideoByID(id string) (*models.Video, error) {
	var p models.Video
	err := repo.Db.Where("ID = ?", id).First(&p).Error
	return &p, err
}

func NewUploadRepository() UploadRepository {
	return &uploadRepository{
		Db: database.Db,
	}
}
