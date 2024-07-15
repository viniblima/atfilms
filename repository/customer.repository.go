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
	RemoveCustomerByID(id string) error
	GetCustomerByID(id string) (*models.Customer, error)
}

func (r *customerRepository) CreateCustomer(c *models.Customer) (*models.Customer, error) {
	err := r.Db.Create(c).Error
	return c, err
}

func (r *customerRepository) ListCustomers() (*[]models.Customer, error) {
	var ls []models.Customer
	err := r.Db.Preload("Logo").Find(&ls).Error
	// err := r.Db.Model(&models.Customer{}).Find(&ls).Error
	return &ls, err
}

func (repo *customerRepository) UpdateCustomer(customer *models.Customer) (*models.Customer, error) {
	err := repo.Db.Save(customer).Error
	return customer, err
}

func (repo *customerRepository) RemoveCustomerByID(id string) error {
	var Customer models.Customer
	err := repo.Db.Model(&models.Customer{}).Where("ID = ?", id).Delete(Customer).Error
	return err
}

func (repo *customerRepository) GetCustomerByID(id string) (*models.Customer, error) {
	var Customer models.Customer
	err := repo.Db.Where("ID = ?", id).Preload("Photo").First(&Customer).Error

	return &Customer, err
}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{
		Db: database.Db,
	}
}
