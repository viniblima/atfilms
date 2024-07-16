package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/controllers"
	"github.com/viniblima/atfilms/middlewares"
)

type UploadRouter interface {
	SetupUploadRoutes(api fiber.Router)
}

type uploadRouter struct {
	uploadController controllers.UploadController
	middleware       middlewares.JWTMiddleware
}

func SetupUploadRoutes(api fiber.Router) {
	router := &uploadRouter{
		uploadController: controllers.NewUploadController(),
		middleware:       middlewares.NewJwtMiddleware(),
	}

	upload_routes := api.Group("/upload")                                                                        // Configuracao da rota pai
	upload_routes.Post("/customer/:id", router.middleware.VerifyJWT, router.uploadController.UploadCustomerLogo) // Upload de foto
}
