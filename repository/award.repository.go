package repository

import (
	"github.com/viniblima/atfilms/database"
	"github.com/viniblima/atfilms/models"
	"gorm.io/gorm"
)

type awardRepository struct {
	Db *gorm.DB
}

type AwardRepository interface {
	Create(a *models.Award) (*models.Award, error)
	GetAll() (*[]models.Award, error)
	GetByID(id string) (*models.Award, error)
	Update(a *models.Award) (*models.Award, error)
	Remove(a *models.Award) error
}

func (r *awardRepository) Create(a *models.Award) (*models.Award, error) {
	err := r.Db.Create(&a).Error
	return a, err
}

func (r *awardRepository) GetAll() (*[]models.Award, error) {
	var ls []models.Award
	err := r.Db.Find(&ls).Error
	return &ls, err
}

func (r *awardRepository) GetByID(id string) (*models.Award, error) {
	var award models.Award
	err := r.Db.Where("ID = ?", id).Preload("AwardImage").Preload("Jobs").First(&award).Error
	return &award, err
}

func (r *awardRepository) Update(a *models.Award) (*models.Award, error) {
	err := r.Db.Omit("Jobs").Save(a).Error
	return a, err
}

func (r *awardRepository) Remove(a *models.Award) error {
	err := r.Db.Delete(&a).Error
	return err
}

func NewAwardRepository() AwardRepository {
	return &awardRepository{
		Db: database.Db,
	}
}
