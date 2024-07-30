package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/controllers"
	"github.com/viniblima/atfilms/middlewares"
)

type TagRouter interface {
	SetupTagRoutes(api fiber.Router)
}

type tagRouter struct {
	tagController controllers.TagController
	middleware    middlewares.JWTMiddleware
}

func SetupTagRoutes(api fiber.Router) {
	router := &tagRouter{
		tagController: controllers.NewTagController(),
		middleware:    middlewares.NewJwtMiddleware(),
	}

	tag_routes := api.Group("/tags")                                              // Configuracao da rota pai
	tag_routes.Get("/", router.middleware.VerifyJWT, router.tagController.GetAll) // Pega lista de tags
}
