package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/viniblima/atfilms/routes"
)

func setupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/api")
	setupV1Routes(api)
}

func setupV1Routes(api fiber.Router) {

	v1 := api.Group("/v1")

	routes.SetupUserRoutes(v1)
	routes.SetupCustomerRoutes(v1)
	routes.SetupJobRoutes(v1)
	routes.SetupUploadRoutes(v1)
}
