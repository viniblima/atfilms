package repository

import (
	"github.com/viniblima/atfilms/database"
	"github.com/viniblima/atfilms/models"
	"gorm.io/gorm"
)

type userRepository struct {
	Db *gorm.DB
}

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

func (repo *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.Db.Where("email = ?", email).First(&user).Error

	return &user, err
}

func (repo *userRepository) CreateUser(user *models.User) (*models.User, error) {
	err := repo.Db.Create(user).Error
	return user, err
}

func NewUserRepository() UserRepository {
	return &userRepository{
		Db: database.Db,
	}
}
