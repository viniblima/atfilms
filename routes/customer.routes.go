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

	customer_routes := api.Group("/customers") // Configuracao da rota pai

	customer_routes.Post("/", router.middleware.VerifyJWT, router.customerController.CreateCustomer)      // Criacao de Customer
	customer_routes.Get("/", router.customerController.ListCustomers)                                     // Lista de Customers
	customer_routes.Get("/:id", router.middleware.VerifyJWT, router.customerController.GetCustomerByID)   // Detalhes de Customer
	customer_routes.Put("/:id", router.middleware.VerifyJWT, router.customerController.UpdateCustomer)    // Atualiza Customer
	customer_routes.Delete("/:id", router.middleware.VerifyJWT, router.customerController.RemoveCustomer) // Remove Customer
}
