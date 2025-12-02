package main

import (
	"absensi-versevox/config"
	"absensi-versevox/database/uow"
	"absensi-versevox/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// inisialisasi konfigurasi dan koneksi database
	config.LoadEnv()
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// Inisialisasi UnitOfWork
	_ = uow.NewUnitOfWork(db)

	// 1. Inisialisasi Fiber
	app := fiber.New()

	// 2. Definisikan Route (Endpoint)
	routes.Setup(app)

	// 3. server berjalan pada port yang ditentukan di konfigurasi
	port := config.AppConfig.AppPort
	log.Println("server is running on port: ", port)
	log.Fatal(app.Listen(":" + port))
}
