package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/repository"
)

type TagController interface {
	GetAll(c *fiber.Ctx) error
}

type tagController struct {
	tagRepo repository.TagRepository
}

func (controller tagController) GetAll(c *fiber.Ctx) error {
	tags, err := controller.tagRepo.GetAll()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(&tags)
}

func NewTagController() TagController {
	return &tagController{
		tagRepo: repository.NewTagRepository(),
	}
}
