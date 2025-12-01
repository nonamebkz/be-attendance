package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Definisikan route di sini
	app.Get("/v1/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
