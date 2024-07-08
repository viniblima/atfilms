package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/controllers"
)

type UserRouter interface {
	SetupUserRoutes(api fiber.Router)
}
type userRouter struct {
	userController controllers.UserController
}

/*
Configura as rotas de usuário
*/
func SetupUserRoutes(api fiber.Router) {
	router := userRouter{}
	router.userController = controllers.NewUserController()
	user_routes := api.Group("/users") // Configuracao da rota pai

	user_routes.Post("/signup", router.userController.SignUp) // Criacao de usuário
	user_routes.Post("/signin", router.userController.SignIn) // Login do usuário
}

// func NewUserRouter() UserRouter {
// 	return &userRouter{}
// }
