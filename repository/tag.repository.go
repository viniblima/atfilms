package repository

import (
	"github.com/viniblima/atfilms/database"
	"github.com/viniblima/atfilms/models"
	"gorm.io/gorm"
)

type tagRepository struct {
	Db *gorm.DB
}

type TagRepository interface {
	Create(m *models.Tag) (*models.Tag, error)
	GetByName(name string) (*models.Tag, error)
	GetAll() (*[]models.Tag, error)
}

func (r *tagRepository) Create(m *models.Tag) (*models.Tag, error) {
	err := r.Db.Omit("Jobs").Create(&m).Error
	return m, err
}

func (r *tagRepository) GetByName(name string) (*models.Tag, error) {
	var tag models.Tag
	err := r.Db.Where("Name = ?", name).First(&tag).Error
	return &tag, err
}

func (r *tagRepository) GetAll() (*[]models.Tag, error) {
	var ls []models.Tag
	err := r.Db.Find(&ls).Error
	return &ls, err
}

func NewTagRepository() TagRepository {
	return &tagRepository{
		Db: database.Db,
	}
}
