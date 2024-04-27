package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lgmztech/profanity-filter/handlers"
)

func SetupRoutes(app *fiber.App) {
	defaultApi(app)
	defaultRoutes(app)
}

func defaultApi(app *fiber.App) {
	app.Get("/v1/:id", handlers.GETProfanityCheck)
}

func defaultRoutes(app *fiber.App) {
	app.Get("/", handlers.GETHomeHandler)
}
