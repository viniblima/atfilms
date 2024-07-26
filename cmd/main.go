package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/database"
)

/*
Inicia o projeto com a seguinte ordem:
1. Inicia o Banco de dados com as devidas migrations
2. Instancia o Fiber
3. Adiciona as rotas das APIs
3. Inicia a subscricao na devida porta
*/
func main() {
	database.ConnectDb()
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	setupRoutes(app)
	app.Static("/upload", "./uploads")
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
