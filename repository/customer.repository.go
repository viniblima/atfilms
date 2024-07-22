package repository

import (
	"github.com/viniblima/atfilms/database"
	"github.com/viniblima/atfilms/models"
	"gorm.io/gorm"
)

type customerRepository struct {
	Db *gorm.DB
}

type CustomerRepository interface {
	CreateCustomer(c *models.Customer) (*models.Customer, error)
	ListCustomers() (*[]models.Customer, error)
	UpdateCustomer(customer *models.Customer) (*models.Customer, error)
	RemoveCustomerByID(*models.Customer) error
	GetCustomerByID(id string) (*models.Customer, error)
	GetCustomersHome() (*[]models.Customer, error)
}

func (r *customerRepository) CreateCustomer(c *models.Customer) (*models.Customer, error) {
	err := r.Db.Omit("Logo").Create(c).Error
	return c, err
}

func (r *customerRepository) ListCustomers() (*[]models.Customer, error) {
	var ls []models.Customer
	err := r.Db.Preload("Logo").Find(&ls).Error
	return &ls, err
}

func (repo *customerRepository) UpdateCustomer(customer *models.Customer) (*models.Customer, error) {
	err := repo.Db.Omit("Logo").Save(customer).Error
	return customer, err
}

func (repo *customerRepository) RemoveCustomerByID(customer *models.Customer) error {
	err := repo.Db.Delete(customer).Error
	return err
}

func (repo *customerRepository) GetCustomerByID(id string) (*models.Customer, error) {
	var Customer models.Customer
	err := repo.Db.Where("ID = ?", id).Preload("Logo").First(&Customer).Error

	return &Customer, err
}

func (repo *customerRepository) GetCustomersHome() (*[]models.Customer, error) {
	var ls []models.Customer
	err := repo.Db.Where("show_in_home = ?", true).Preload("Logo").Find(&ls).Error
	return &ls, err
}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{
		Db: database.Db,
	}
}
