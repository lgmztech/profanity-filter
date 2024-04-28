package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lgmztech/profanity-filter/handlers"
)

func SetupRoutes(app *fiber.App) {
	defaultApi(app)
	helperRoutes(app)
}

func defaultApi(app *fiber.App) {
	app.Get("/v1/profanityfilter/:id", handlers.GETProfanityCheck)
}

func helperRoutes(app *fiber.App) {
	app.Get("/", handlers.GETProfanityFilterHelpHandler)
}
