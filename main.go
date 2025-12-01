package main

import (
	"absensi-versevox/config"
	"absensi-versevox/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// inisialisasi konfigurasi dan koneksi database
	config.LoadEnv()
	config.ConnectDB()

	// 1. Inisialisasi Fiber
	app := fiber.New()

	// 2. Definisikan Route (Endpoint)
	// Ketika user mengakses path "/" dengan metode GET
	routes.Setup(app)

	// server berjalan pada port yang ditentukan di konfigurasi
	port := config.AppConfig.AppPort
	log.Println("server is running on port: ", port)
	log.Fatal(app.Listen(":" + port))
}
