package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/handlers"
	"github.com/viniblima/atfilms/models"
	"github.com/viniblima/atfilms/repository"
)

type ClientController interface {
	CreateClient(ctx *fiber.Ctx) error
	UpdateClient(ctx *fiber.Ctx) error
	ListClients(ctx *fiber.Ctx) error
	GetClientByID(c *fiber.Ctx) error
}

type clientController struct {
	clientRepo repository.ClientRepository
}

type CreateClientStruct struct {
	Name string `json:"Name" validate:"required,min=3,max=32"`
}

func (controller *clientController) CreateClient(c *fiber.Ctx) error {
	body := new(CreateClientStruct)
	c.BodyParser(&body)
	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	client := models.Client{
		Name: body.Name,
	}
	newClient, err := controller.clientRepo.CreateClient(&client)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.Status(http.StatusCreated).JSON(newClient)
}

func (controller *clientController) GetClientByID(c *fiber.Ctx) error {
	id := c.Params("id")
	client, errClient := controller.clientRepo.GetClientByID(id)

	if errClient != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Client not found",
		})
	}
	return c.Status(http.StatusOK).JSON(client)
}

func (controller *clientController) UpdateClient(c *fiber.Ctx) error {
	body := new(CreateClientStruct)
	c.BodyParser(&body)
	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}

	id := c.Params("id")

	client, errClient := controller.clientRepo.GetClientByID(id)

	if errClient != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Client not found",
		})
	}

	client.Name = body.Name

	controller.clientRepo.UpdateClient(client)

	return c.Status(http.StatusOK).JSON(fiber.Map{

		"Client": *client,
	})
}

func (controller *clientController) ListClients(c *fiber.Ctx) error {
	clients, err := controller.clientRepo.ListClients()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(&clients)
}

func NewClientController() ClientController {
	return &clientController{}
}
