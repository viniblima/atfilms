package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/controllers"
)

type HomeRouter interface {
	SetupHomeRoutes(api fiber.Router)
}

type homeRouter struct {
	homeController controllers.HomeController
}

func SetupHomeRoutes(api fiber.Router) {
	router := &homeRouter{
		homeController: controllers.NewHomeController(),
	}

	home_router := api.Group("/home")                   // Configuracao da rota pai
	home_router.Get("/", router.homeController.GetHome) // Pega lista de cases e clientes
}
