package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/handlers"
	"github.com/viniblima/atfilms/models"
	"github.com/viniblima/atfilms/repository"
)

type AwardController interface {
	CreateAward(c *fiber.Ctx) error
	ListAwards(c *fiber.Ctx) error
	GetAwardByID(c *fiber.Ctx) error
	UpdateAward(c *fiber.Ctx) error
	DeleteAward(c *fiber.Ctx) error
}

type awardController struct {
	awardRepo repository.AwardRepository
}

type CreateAwardStruct struct {
	Name string `json:"Name" validate:"required,min=3,max=32"`
}

func (controller awardController) CreateAward(c *fiber.Ctx) error {
	body := new(CreateAwardStruct)
	c.BodyParser(&body)
	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}

	award := models.Award{
		Name: body.Name,
	}

	newAward, err := controller.awardRepo.Create(&award)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.Status(http.StatusCreated).JSON(newAward)
}

func (controller awardController) ListAwards(c *fiber.Ctx) error {
	awards, err := controller.awardRepo.GetAll()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(&awards)
}

func (controller awardController) GetAwardByID(c *fiber.Ctx) error {
	id := c.Params("id")

	award, err := controller.awardRepo.GetByID(id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Award not found",
		})
	}
	return c.Status(http.StatusOK).JSON(award)
}

func (controller awardController) UpdateAward(c *fiber.Ctx) error {
	body := new(CreateAwardStruct)
	c.BodyParser(&body)
	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	id := c.Params("id")

	award, errAward := controller.awardRepo.GetByID(id)

	if errAward != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Award not found",
		})
	}
	award.Name = body.Name

	update, errUpdate := controller.awardRepo.Update(award)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpdate)
	}

	return c.Status(http.StatusOK).JSON(update)
}

func (controller awardController) DeleteAward(c *fiber.Ctx) error {
	id := c.Params("id")

	award, errAward := controller.awardRepo.GetByID(id)

	if errAward != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Award not found",
		})
	}

	resultRemove, errRemove := handlers.RemoveS3(award.AwardImage.FileName)

	if errRemove != nil {
		return c.Status(http.StatusBadRequest).JSON(errRemove)
	}

	errR := controller.awardRepo.Remove(award)

	if errR != nil {
		return c.Status(http.StatusBadRequest).JSON(errR)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Msg":    "Award removed",
		"Result": resultRemove,
	})
}

func NewAwardController() AwardController {
	return &awardController{
		awardRepo: repository.NewAwardRepository(),
	}
}
