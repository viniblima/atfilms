package repository

import (
	"github.com/viniblima/atfilms/database"
	"github.com/viniblima/atfilms/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	AppendTag(job *models.Job, tag []*models.Tag) (*models.Job, error, error)
	AppendAward(job *models.Job, awards []*models.Award) (*models.Job, error, error)
}

func (r *jobRepository) CreateJob(m *models.Job) (*models.Job, error) {
	err := r.Db.Omit("Placeholder", "Video", "Customer", "Components", "Tags", "Awards").Create(&m).Error
	return m, err
}

func (r *jobRepository) ListAllJobs() (*[]models.Job, error) {
	var ls []models.Job
	err := r.Db.Preload("Customer").Preload("Components").Preload("Placeholder").Preload("MainVideo").Find(&ls).Error
	return &ls, err
}

func (r *jobRepository) UpdateJob(job *models.Job) (*models.Job, error) {
	err := r.Db.Omit("Placeholder", "Video", "Customer", "Components", "Tags", "Awards").Save(&job).Error
	return job, err
}

func (repo *jobRepository) GetJobByID(id string) (*models.Job, error) {
	var job models.Job
	err := repo.Db.Preload("Awards."+clause.Associations).Preload("Customer").Preload("Components."+clause.Associations).Preload("Placeholder").Preload("MainVideo").Preload("Tags").Where("ID = ?", id).Or("Slug = ?", id).First(&job).Error

	return &job, err
}

func (repo *jobRepository) GetJobsHome() (*[]models.Job, error) {
	var ls []models.Job
	err := repo.Db.Where("show_in_home = ?", true).Preload("Customer").Preload("Components").Preload("Placeholder").Preload("MainVideo").Find(&ls).Error
	return &ls, err
}

func (repo *jobRepository) RemoveJobByID(job *models.Job) error {
	err := repo.Db.Delete(&job).Error
	return err
}

func (repo *jobRepository) AppendTag(job *models.Job, tag []*models.Tag) (*models.Job, error, error) {
	errClear := repo.Db.Model(&job).Association("Tags").Clear()
	err := repo.Db.Model(&job).Omit("Tags.*").Association("Tags").Append(&tag)
	return job, err, errClear
}

func (repo *jobRepository) AppendAward(job *models.Job, awards []*models.Award) (*models.Job, error, error) {
	errClear := repo.Db.Model(&job).Association("Awards").Clear()
	err := repo.Db.Model(&job).Omit("Awards.*").Association("Awards").Append(&awards)

	return job, err, errClear
}

func NewJobRepository() JobRepository {
	return &jobRepository{
		Db: database.Db,
	}
}
