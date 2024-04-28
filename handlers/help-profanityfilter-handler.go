package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lgmztech/profanity-filter/database"
	"github.com/lgmztech/profanity-filter/middleware"
)

func GETProfanityFilterHelpHandler(c *fiber.Ctx) error {
	cookie := c.Cookies("auth_token")
	validatedCookie := middleware.ValidateJWT(cookie)

	if validatedCookie == "" {
		userUUID := uuid.New()
		newJWTToken := middleware.GenerateJWT(userUUID.String())
		database.CreateNewToken(userUUID.String())
		fmt.Println(userUUID.String())

		cookie := fiber.Cookie{
			Name:  "auth_token",
			Value: newJWTToken,
		}

		c.Cookie(&cookie)
		validatedCookie = userUUID.String()
	}

	return c.Render("help-profanityfilter", fiber.Map{
		"UUID": validatedCookie,
	})
}
