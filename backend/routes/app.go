package routes

import (
	"i9codesauths/backend/helpers"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func App(router fiber.Router) {
	router.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("AUTH_JWT_SECRET"))},
		ContextKey: "auth",
	}))

	// access a restricted resource : jwt auth
	router.Get("/restricted", func(c *fiber.Ctx) error {
		type User struct {
			Id       int    `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		}

		userMap := c.Locals("auth").(*jwt.Token).Claims.(jwt.MapClaims)["data"].(map[string]any)

		var user User

		helpers.MapToStruct(userMap, &user)

		return c.JSON(user)
	})
}
