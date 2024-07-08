package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/viniblima/atfilms/repository"
)

type JWTMiddleware interface {
	VerifyJWT(c *fiber.Ctx) error
}
type jwtMiddleware struct {
	userRepo repository.UserRepository
}

/*
Funcao que verifica o JWT enviado pelo client. A funcao se baseia no seguintes passos:
 1. Verificacao se o cabecalho da requisicao está correto; caso haja erro, retorna resposta com erro com o status 401;
 2. Tentativa de extrair e validar o JWT enviado; caso haja erro, retorna resposta com erro com o status 401
 3. Insercao de ID do usuário nos Locals para acesso à esse dado posteriormente
    e para evitar que, mesmo que o JWT esteja correto,
    haja um contorno e tentativa de uma requisicao ese passando por outro usuário
*/
func (controller *jwtMiddleware) VerifyJWT(c *fiber.Ctx) error {

	auth := c.Get("Authorization")
	claims := jwt.MapClaims{}
	if auth != "" {
		split := strings.Split(auth, "JWT ")

		if len(split) < 2 {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid tag",
			})
		}

		_, err := jwt.ParseWithClaims(split[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("PASSWORD_SECRET")), nil
		})

		if claims["sub"] != "auth" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": err.Error(),
			})

		}

	} else {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid headers",
		})
	}

	c.Locals("userID", claims["id"])
	c.Next()
	return nil
}

func NewJwtMiddleware() JWTMiddleware {
	return &jwtMiddleware{}
}
