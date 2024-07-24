package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/handlers"
	"github.com/viniblima/atfilms/models"
	"github.com/viniblima/atfilms/repository"
)

type JobController interface {
	CreateJob(c *fiber.Ctx) error
	UpdateJob(c *fiber.Ctx) error
	ListAllJobs(c *fiber.Ctx) error
	GetJobBySlug(c *fiber.Ctx) error
}

type jobController struct {
	jobRepo      repository.JobRepository
	customerRepo repository.CustomerRepository
}

type CreateJobStruct struct {
	Name       string                `json:"Name" validate:"required,min=3,max=32"`
	Slug       string                `json:"Slug" validate:"required,min=3,max=32"`
	CustomerID string                `json:"CustomerID" validate:"required"`
	ShowInHome bool                  `json:"ShowInHome" validate:"required"`
	Components []models.JobComponent `json:"Components"`
}

func (controller jobController) UpdateJob(c *fiber.Ctx) error {
	body := new(CreateJobStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	id := c.Params("id")

	jobFound, errJobFound := controller.jobRepo.GetJobByID(id)

	if errJobFound != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Job not found",
		})
	}

	Customer, errCustomer := controller.customerRepo.GetCustomerByID(body.CustomerID)

	if errCustomer != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	jobFound.Customer = *Customer
	jobFound.Name = body.Name
	jobFound.Slug = body.Slug
	jobFound.ShowInHome = body.ShowInHome

	update, errUpdate := controller.jobRepo.UpdateJob(jobFound)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpdate)
	}

	return c.Status(http.StatusOK).JSON(update)
}

func (controller jobController) CreateJob(c *fiber.Ctx) error {
	body := new(CreateJobStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}

	customer, errCustomer := controller.customerRepo.GetCustomerByID(body.CustomerID)

	if errCustomer != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	newJob := models.Job{
		Name:       body.Name,
		Slug:       body.Slug,
		ShowInHome: body.ShowInHome,
		Components: body.Components,
	}
	newJob.Customer = *customer

	job, errJob := controller.jobRepo.CreateJob(&newJob)

	if errJob != nil {
		return c.Status(http.StatusBadRequest).JSON(errJob)
	}

	return c.Status(http.StatusCreated).JSON(job)
}

func (controller jobController) ListAllJobs(c *fiber.Ctx) error {
	jobs, err := controller.jobRepo.ListAllJobs()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(&jobs)
}

func (controller jobController) GetJobBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	job, err := controller.jobRepo.GetJobBySlug(slug)

	if err != nil {

		job, err := controller.jobRepo.GetJobByID(slug)

		if err != nil {

			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Job not found",
			})
		}
		return c.Status(http.StatusOK).JSON(job)
	}
	return c.Status(http.StatusOK).JSON(job)
}

func NewJobController() JobController {
	return &jobController{
		jobRepo: repository.NewJobRepository(),
	}
}
