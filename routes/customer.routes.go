package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/controllers"
	"github.com/viniblima/atfilms/middlewares"
)

type CustomerRouter interface {
	SetupCustomerRoutes(api fiber.Router)
}

type customerRouter struct {
	customerController controllers.CustomerController
	middleware         middlewares.JWTMiddleware
}

func SetupCustomerRoutes(api fiber.Router) {
	router := &customerRouter{
		customerController: controllers.NewCustomerController(),
		middleware:         middlewares.NewJwtMiddleware(),
	}

	Customer_routes := api.Group("/customers") // Configuracao da rota pai

	Customer_routes.Post("/", router.middleware.VerifyJWT, router.customerController.CreateCustomer) // Criacao de Customer
	Customer_routes.Get("/", router.middleware.VerifyJWT, router.customerController.ListCustomers)   // Lista de Customers

	Customer_routes.Get("/:id", router.middleware.VerifyJWT, router.customerController.GetCustomerByID) // Detalhes de Customer
	Customer_routes.Put("/:id", router.middleware.VerifyJWT, router.customerController.UpdateCustomer)  // Atualiza Customer

	Customer_routes.Delete("/:id", router.middleware.VerifyJWT, router.customerController.RemoveCustomer) // Remove Customer
}
