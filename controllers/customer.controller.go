package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/handlers"
	"github.com/viniblima/atfilms/models"
	"github.com/viniblima/atfilms/repository"
)

type CustomerController interface {
	CreateCustomer(ctx *fiber.Ctx) error
	UpdateCustomer(ctx *fiber.Ctx) error
	ListCustomers(ctx *fiber.Ctx) error
	GetCustomerByID(c *fiber.Ctx) error
	RemoveCustomer(c *fiber.Ctx) error
}

type customerController struct {
	customerRepo repository.CustomerRepository
}

type CreateCustomerStruct struct {
	Name       string `json:"Name" validate:"required,min=3,max=32"`
	ShowInHome bool   `json:"ShowInHome"`
	Position   int    `json:"Position" validate:"required"`
}

func (controller customerController) CreateCustomer(c *fiber.Ctx) error {
	body := new(CreateCustomerStruct)
	c.BodyParser(&body)
	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	ls, errLs := controller.customerRepo.ListCustomers()

	if errLs != nil {
		return c.Status(http.StatusBadRequest).JSON(errLs)
	}

	customer := models.Customer{
		Name:       body.Name,
		ShowInHome: &body.ShowInHome,
		Position:   len(*ls),
	}
	newCustomer, err := controller.customerRepo.CreateCustomer(&customer)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.Status(http.StatusCreated).JSON(newCustomer)
}

func (controller customerController) GetCustomerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	customer, errCustomer := controller.customerRepo.GetCustomerByID(id)

	if errCustomer != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}
	return c.Status(http.StatusOK).JSON(customer)
}

func (controller customerController) UpdateCustomer(c *fiber.Ctx) error {
	body := new(CreateCustomerStruct)
	c.BodyParser(&body)
	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}

	id := c.Params("id")

	customer, errCustomer := controller.customerRepo.GetCustomerByID(id)

	if errCustomer != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	customer.Name = body.Name
	customer.ShowInHome = &body.ShowInHome
	customer.Position = body.Position

	controller.customerRepo.UpdateCustomer(customer)

	return c.Status(http.StatusOK).JSON(fiber.Map{

		"Customer": *customer,
	})
}

func (controller customerController) ListCustomers(c *fiber.Ctx) error {
	customers, err := controller.customerRepo.ListCustomers()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(&customers)
}

func (controller customerController) RemoveCustomer(c *fiber.Ctx) error {

	id := c.Params("id")

	customer, errCustomer := controller.customerRepo.GetCustomerByID(id)

	if errCustomer != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	folder := fmt.Sprintf("tmp/uploads/%s", customer.Logo.FileName)

	os.Remove(folder)

	errR := controller.customerRepo.RemoveCustomerByID(customer)

	if errR != nil {
		return c.Status(http.StatusBadRequest).JSON(errR)
	}

	ls, errLs := controller.customerRepo.ListCustomers()

	if errLs != nil {
		return c.Status(http.StatusBadRequest).JSON(errLs)
	}

	for i := 0; i < len(*ls); i++ {
		newCs := (*ls)[i]
		newCs.Position = i
		controller.customerRepo.UpdateCustomer(&newCs)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Msg": "Customer removed",
	})
}

func NewCustomerController() CustomerController {
	return &customerController{
		customerRepo: repository.NewCustomerRepository(),
	}
}
