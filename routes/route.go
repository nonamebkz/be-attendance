package routes

import (
	"absensi-versevox/database/uow"
	"absensi-versevox/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App, uowInstance uow.UnitOfWork) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Definisikan route di sini
	app.Get("/v1/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Karyawan endpoints
	app.Get("/karyawan-get", func(c *fiber.Ctx) error {
		return handlers.GetKaryawanByID(c, uowInstance)
	})
	app.Post("/karyawan-create", func(c *fiber.Ctx) error {
		return handlers.CreateKaryawanByID(c, uowInstance)
	})
	app.Put("/karyawan-update", func(c *fiber.Ctx) error {
		return handlers.UpdateKaryawanByID(c, uowInstance)
	})
	app.Delete("/karyawan-delete", func(c *fiber.Ctx) error {
		return handlers.DeleteKaryawanByID(c, uowInstance)
	})

	//ini contoh ya
	app.Post("/karyawan-create-tx", func(c *fiber.Ctx) error {
		return handlers.CreateKaryawanWithTX(c, uowInstance)
	})

	//Pengguna endpoints
	app.Get("/pengguna-get", func(c *fiber.Ctx) error {
		return handlers.GetPenggunaByID(c, uowInstance)
	})
	app.Post("/pengguna-create", func(c *fiber.Ctx) error {
		return handlers.CreatePenggunaByID(c, uowInstance)
	})
}
