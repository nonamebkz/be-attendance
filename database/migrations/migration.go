package migrations

import (
	"absensi-versevox/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	fmt.Println("⏳ Memulai migrasi database...")

	err := db.AutoMigrate(
		&models.Karyawan{},
	)

	if err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}

	fmt.Println("✅ Migrasi database selesai.")

}
