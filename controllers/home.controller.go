package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/repository"
)

type HomeController interface {
	GetHome(c *fiber.Ctx) error
}

type homeController struct {
	customerRepo repository.CustomerRepository
	jobRepo      repository.JobRepository
}

func (controller homeController) GetHome(c *fiber.Ctx) error {
	cs, errCs := controller.customerRepo.GetCustomersHome()

	if errCs != nil {
		return c.Status(http.StatusBadRequest).JSON(errCs)
	}

	js, errJs := controller.jobRepo.GetJobsHome()

	if errJs != nil {
		return c.Status(http.StatusBadRequest).JSON(errJs)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"jobs":      js,
		"customers": cs,
	})
}

func NewHomeController() HomeController {
	return &homeController{
		customerRepo: repository.NewCustomerRepository(),
		jobRepo:      repository.NewJobRepository(),
	}
}
