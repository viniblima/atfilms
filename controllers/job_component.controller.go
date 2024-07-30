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
	Remove(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

type jobComponentController struct {
	jobComponentRepo repository.JobComponentRepository
}

func (controller jobComponentController) Create(c *fiber.Ctx) error {

	id := c.Params("id")

	body := new(models.CreateJobComponentStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	j := models.JobComponent{
		Type:  body.Type,
		JobID: id,
	}
	jc, errJc := controller.jobComponentRepo.CreateJobComponent(&j)

	if errJc != nil {
		return c.Status(http.StatusBadRequest).JSON(errJc)
	}

	return c.Status(http.StatusCreated).JSON(jc)
}

func (controller jobComponentController) Remove(c *fiber.Ctx) error {
	id := c.Params("id")

	cpnt, errCpnt := controller.jobComponentRepo.GetJobComponentByID(id)

	if errCpnt != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Job not found",
		})
	}

	if cpnt.Slider != nil {
		for i := 0; i < len(*cpnt.Slider); i++ {
			photo := (*cpnt.Slider)[i]

			_, errRemove := handlers.RemoveS3(photo.FileName)
			if errRemove != nil {
				return c.Status(http.StatusBadRequest).JSON(errRemove)
			}
		}
	}

	if cpnt.Videos != nil {
		for i := 0; i < len(*cpnt.Videos); i++ {
			photo := (*cpnt.Videos)[i]

			_, errRemove := handlers.RemoveS3(photo.FileName)
			if errRemove != nil {
				return c.Status(http.StatusBadRequest).JSON(errRemove)
			}
		}
	}

	if cpnt.FillPhotoHorizontal != nil {
		_, errRemove := handlers.RemoveS3(cpnt.FillPhotoHorizontal.FileName)
		if errRemove != nil {
			return c.Status(http.StatusBadRequest).JSON(errRemove)
		}
	}

	errRemove := controller.jobComponentRepo.Remove(cpnt)

	if errRemove != nil {
		return c.Status(http.StatusBadRequest).JSON(errRemove)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Msg": "JobComponent Removed"})
}

func (controller jobComponentController) Update(c *fiber.Ctx) error {
	body := new(models.CreateJobComponentStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}

	id := c.Params("id")

	cpnt, errCpnt := controller.jobComponentRepo.GetJobComponentByID(id)

	if errCpnt != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Component Job not found",
		})
	}

	cpnt.Position = body.Position
	cpnt.Type = body.Type

	update, errUpdate := controller.jobComponentRepo.UpdateJobComponent(cpnt)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpdate)
	}

	return c.Status(http.StatusOK).JSON(update)
}

func NewJobComponentController() JobComponentController {
	return &jobComponentController{
		jobComponentRepo: repository.NewJobComponentRepository(),
	}
}
