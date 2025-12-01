package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Inisialisasi Fiber
	app := fiber.New()

	// 2. Definisikan Route (Endpoint)
	// Ketika user mengakses path "/" dengan metode GET
	app.Get("/", func(c *fiber.Ctx) error {
		// Mengembalikan response JSON
		return c.JSON(fiber.Map{
			"message": "Hello World! Ini adalah API Absensi Anda.",
			"status":  "success",
		})

		// Atau hanya mengembalikan teks:
		// return c.SendString("Hello, World 👋!")
	})

	// 3. Menjalankan Server
	// Server akan berjalan di port 3000
	log.Fatal(app.Listen(":3000"))
}
