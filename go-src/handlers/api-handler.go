package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lgmztech/profanity-filter/database"
	"github.com/lgmztech/profanity-filter/middleware"
)

func GETProfanityCheck(c *fiber.Ctx) error {
	paramStringToken := c.Params("id")
	if paramStringToken == "" {
		c.Redirect("/")
	}

	_, limitCheck := database.CheckForValidToken(paramStringToken)
	if limitCheck != nil {
		return c.JSON(fiber.Map{
			"error": "API Token Limit Reached",
		})
	}

	// Parse request body
	userMessage := c.FormValue("userMessage")

	// Check message for profanity
	profanityModified, hasProfanity := middleware.CheckMessage(userMessage)

	database.AppendTokenUse(paramStringToken)

	// Return success response
	return c.JSON(fiber.Map{
		"original-message": userMessage,
		"modified-message": profanityModified,
		"has-profanity":    hasProfanity,
	})
}
