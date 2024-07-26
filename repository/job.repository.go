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
	GetJobsHome() (*[]models.Job, error)
	RemoveJobByID(customer *models.Job) error
}

func (r *jobRepository) CreateJob(m *models.Job) (*models.Job, error) {
	err := r.Db.Omit("Placeholder", "Video", "Customer", "Components").Create(&m).Error
	return m, err
}

func (r *jobRepository) ListAllJobs() (*[]models.Job, error) {
	var ls []models.Job
	err := r.Db.Preload("Customer").Preload("Components").Preload("Placeholder").Preload("MainVideo").Find(&ls).Error
	return &ls, err
}

func (r *jobRepository) UpdateJob(job *models.Job) (*models.Job, error) {
	err := r.Db.Omit("Placeholder", "Video", "Customer", "Components").Save(&job).Error
	return job, err
}

func (repo *jobRepository) GetJobByID(id string) (*models.Job, error) {
	var job models.Job
	err := repo.Db.Preload("Customer").Preload("Components").Preload("Placeholder").Preload("MainVideo").Where("ID = ?", id).Or("Slug = ?", id).First(&job).Error

	return &job, err
}

func (repo *jobRepository) GetJobsHome() (*[]models.Job, error) {
	var ls []models.Job
	err := repo.Db.Where("show_in_home = ?", true).Preload("Customer").Preload("Components").Preload("Placeholder").Preload("MainVideo").Find(&ls).Error
	return &ls, err
}

func (repo *jobRepository) RemoveJobByID(customer *models.Job) error {
	err := repo.Db.Delete(&customer).Error
	return err
}

func NewJobRepository() JobRepository {
	return &jobRepository{
		Db: database.Db,
	}
}
