package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/handlers"
	"github.com/viniblima/atfilms/models"
	"github.com/viniblima/atfilms/repository"
)

type UserController interface {
	SignIn(c *fiber.Ctx) error
}

type userController struct {
	userRepo repository.UserRepository
}

func extractUserObj(u models.User) map[string]interface{} {
	return map[string]interface{}{
		"ID":    u.ID,
		"Name":  u.Name,
		"Email": u.Email,
	}
}

type LoginStruct struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required"`
}

func (controller userController) SignIn(c *fiber.Ctx) error {

	body := new(LoginStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	print(body.Email)
	u, errorU := controller.userRepo.GetUserByEmail(body.Email)
	print(u)

	checked := handlers.CheckHash(u.Password, body.Password)

	if errorU != nil || !checked {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Email or password wrong"})
	}

	json, errJwt := handlers.GenerateJWT(u.ID)

	if errJwt != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Error on generate JWT"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Auth": json,
		"User": extractUserObj(*u),
	})

}

func NewUserController() UserController {
	return &userController{
		userRepo: repository.NewUserRepository(),
	}
}
