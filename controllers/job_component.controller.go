package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/handlers"
	"github.com/viniblima/atfilms/models"
	"github.com/viniblima/atfilms/repository"
)

type JobComponentController interface {
	Create(c *fiber.Ctx) error
}

type jobComponentController struct {
	jobComponentRepo repository.JobComponentRepository
}

func (controller jobComponentController) Create(c *fiber.Ctx) error {
	body := new(models.CreateJobComponentStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	j := models.JobComponent{
		Type: body.Type,
	}
	jc, errJc := controller.jobComponentRepo.CreateJobComponent(&j)

	if errJc != nil {
		return c.Status(http.StatusBadRequest).JSON(errJc)
	}

	return c.Status(http.StatusCreated).JSON(jc)
}

func NewJobComponentController() JobComponentController {
	return &jobComponentController{
		jobComponentRepo: repository.NewJobComponentRepository(),
	}
}
