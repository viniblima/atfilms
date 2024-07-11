package repository

import (
	"github.com/viniblima/atfilms/database"
	"github.com/viniblima/atfilms/models"
	"gorm.io/gorm"
)

type clientRepository struct {
	Db *gorm.DB
}

type ClientRepository interface {
	CreateClient(c *models.Client) (*models.Client, error)
	ListClients() (*[]models.Client, error)
	UpdateClient(client *models.Client) (*models.Client, error)
	RemoveClientByID(id string) error
	GetClientByID(id string) (*models.Client, error)
}

func (r *clientRepository) CreateClient(c *models.Client) (*models.Client, error) {
	err := r.Db.Create(c).Error
	return c, err
}

func (r *clientRepository) ListClients() (*[]models.Client, error) {
	var ls []models.Client
	err := r.Db.Find(&ls).Error
	return &ls, err
}

func (repo *clientRepository) UpdateClient(client *models.Client) (*models.Client, error) {
	err := repo.Db.Save(client).Error
	return client, err
}

func (repo *clientRepository) RemoveClientByID(id string) error {
	var client models.Client
	err := repo.Db.Model(&models.Client{}).Where("ID = ?", id).Delete(client).Error
	return err
}

func (repo *clientRepository) GetClientByID(id string) (*models.Client, error) {
	var client models.Client
	err := repo.Db.Where("ID = ?", id).First(&client).Error

	return &client, err
}

func NewClientRepository() ClientRepository {
	return &clientRepository{
		Db: database.Db,
	}
}
