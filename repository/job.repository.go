package repository

import (
	"github.com/viniblima/atfilms/database"
	"github.com/viniblima/atfilms/models"
	"gorm.io/gorm"
)

type jobRepository struct {
	Db *gorm.DB
}

type JobRepository interface {
	CreateJob(m *models.Job) (*models.Job, error)
	ListAllJobs() (*[]models.Job, error)
	UpdateJob(m *models.Job) (*models.Job, error)
	GetJobByID(id string) (*models.Job, error)
	GetJobBySlug(slug string) (*models.Job, error)
}

func (r *jobRepository) CreateJob(m *models.Job) (*models.Job, error) {
	err := r.Db.Create(m).Error
	return m, err
}

func (r *jobRepository) ListAllJobs() (*[]models.Job, error) {
	var ls []models.Job
	err := r.Db.Find(&ls).Error
	return &ls, err
}

func (r *jobRepository) UpdateJob(job *models.Job) (*models.Job, error) {
	err := r.Db.Save(job).Error
	return job, err
}

func (repo *jobRepository) GetJobByID(id string) (*models.Job, error) {
	var job models.Job
	err := repo.Db.Where("ID = ?", id).First(&job).Error

	return &job, err
}

func (repo *jobRepository) GetJobBySlug(slug string) (*models.Job, error) {
	var job models.Job
	err := repo.Db.Where("Slug = ?", slug).First(&job).Error

	return &job, err
}

func NewJobRepository() JobRepository {
	return &jobRepository{
		Db: database.Db,
	}
}
