package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/controllers"
	"github.com/viniblima/atfilms/middlewares"
)

type ClientRouter interface {
	SetupClientRoutes(api fiber.Router)
}

type clientRouter struct {
	clientController controllers.ClientController
	middleware       middlewares.JWTMiddleware
}

func SetupClientRoutes(api fiber.Router) {
	router := &clientRouter{
		clientController: controllers.NewClientController(),
		middleware:       middlewares.NewJwtMiddleware(),
	}

	client_routes := api.Group("/clients") // Configuracao da rota pai

	client_routes.Post("/", router.middleware.VerifyJWT, router.clientController.CreateClient) // Criacao de cliente
	client_routes.Get("/", router.middleware.VerifyJWT, router.clientController.ListClients)   // Lista de clientes

	client_routes.Get("/:id", router.middleware.VerifyJWT, router.clientController.GetClientByID) // Detalhes de cliente
	client_routes.Put("/:id", router.middleware.VerifyJWT, router.clientController.UpdateClient)  // Atualiza cliente
}
