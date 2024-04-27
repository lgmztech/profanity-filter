package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lgmztech/profanity-filter/database"
	"github.com/lgmztech/profanity-filter/middleware"
)

func GETHomeHandler(c *fiber.Ctx) error {
	var userUUID string
	cookie := c.Cookies("auth_token")

	fmt.Println(cookie)

	validatedCookie := middleware.ValidateJWT(cookie)

	fmt.Println(validatedCookie)

	if validatedCookie == "" {
		userUUID := uuid.New()
		middleware.GenerateJWT(userUUID.String())
		database.CreateNewToken(userUUID.String())
		fmt.Println(userUUID.String())

		cookie := fiber.Cookie{
			Name:  "auth_token",
			Value: validatedCookie,
		}

		c.Cookie(&cookie)
	}

	return c.JSON(fiber.Map{
		"user-token": userUUID,
	})
}
