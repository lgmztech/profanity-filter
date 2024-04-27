package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/lgmztech/profanity-filter/database"
	"github.com/lgmztech/profanity-filter/env"
	"github.com/lgmztech/profanity-filter/router"
)

func init() {
	// Load environment variables
	env.LoadEnv()
	// Connect to database
	database.ConnectToDatabase()
	// Sync database
	database.SyncDB()
}

func main() {
	engine := django.New("./views", ".django")
	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/public", "./public")

	router.SetupRoutes(app)

	app.Listen(":3000")

}
