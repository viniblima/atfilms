package repository

import (
	"github.com/viniblima/atfilms/database"
	"github.com/viniblima/atfilms/models"
	"gorm.io/gorm"
)

type jobComponentRepository struct {
	Db *gorm.DB
}

type JobComponentRepository interface {
	CreateJobComponent(m *models.JobComponent) (*models.JobComponent, error)
	UpdateJobComponent(m *models.JobComponent) (*models.JobComponent, error)
	GetJobComponentByID(id string) (*models.JobComponent, error)
	AppendPhotoToSlider(m *models.JobComponent, p *models.Photo) (*models.JobComponent, error)
	AppendVideo(m *models.JobComponent, p *models.Video) (*models.JobComponent, error)
}

func (r *jobComponentRepository) CreateJobComponent(m *models.JobComponent) (*models.JobComponent, error) {
	err := r.Db.Create(m).Error
	return m, err
}

func (r *jobComponentRepository) UpdateJobComponent(m *models.JobComponent) (*models.JobComponent, error) {
	err := r.Db.Save(m).Error
	return m, err
}

func (r *jobComponentRepository) GetJobComponentByID(id string) (*models.JobComponent, error) {
	var jc models.JobComponent
	err := r.Db.Where("ID = ?", id).First(&jc).Error

	return &jc, err
}

func (r *jobComponentRepository) AppendPhotoToSlider(m *models.JobComponent, p *models.Photo) (*models.JobComponent, error) {
	err := r.Db.Model(&m).Omit("Slider.*").Association("Slider").Append(&p)
	return m, err
}

func (r *jobComponentRepository) AppendVideo(m *models.JobComponent, p *models.Video) (*models.JobComponent, error) {
	err := r.Db.Model(&m).Omit("Videos.*").Association("Videos").Append(&p)
	return m, err
}

func NewJobComponentRepository() JobComponentRepository {
	return &jobComponentRepository{
		Db: database.Db,
	}
}
