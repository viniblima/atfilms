package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/controllers"
	"github.com/viniblima/atfilms/middlewares"
)

type AwardRouter interface {
	SetupAwardRoutes(api fiber.Router)
}

type awardRouter struct {
	awardController controllers.AwardController
	middleware      middlewares.JWTMiddleware
}

func SetupAwardRoutes(api fiber.Router) {
	router := &awardRouter{
		awardController: controllers.NewAwardController(),
		middleware:      middlewares.NewJwtMiddleware(),
	}

	award_routes := api.Group("/awards") // Configuracao da rota pai

	award_routes.Post("/", router.middleware.VerifyJWT, router.awardController.CreateAward)      // Criação de prêmio
	award_routes.Get("/", router.awardController.ListAwards)                                     // Lista de prêmios
	award_routes.Get("/:id", router.awardController.GetAwardByID)                                // Detalhe de um prêmio
	award_routes.Put("/:id", router.middleware.VerifyJWT, router.awardController.UpdateAward)    // Atualiza um prêmio
	award_routes.Delete("/:id", router.middleware.VerifyJWT, router.awardController.DeleteAward) // Remove um prêmio
}
