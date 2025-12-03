package handlers

import (
	"absensi-versevox/database/uow"
	"absensi-versevox/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetPenggunaByID handles GET /pengguna-get endpoint
func GetPenggunaByID(c *fiber.Ctx, uowInstance uow.UnitOfWork) error {
	// Get id from query parameter
	idParam := c.Query("id")
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id parameter is required",
		})
	}

	// Convert string to int32
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id format",
		})
	}

	// Get repository from UOW (no transaction needed for read operation)
	penggunaRepo := uowInstance.PenggunaRepository()

	// Get karyawan by ID
	pengguna, err := penggunaRepo.GetPenggunaByID(int32(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "karyawan not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    pengguna,
	})
}

func CreatePenggunaByID(c *fiber.Ctx, uowInstance uow.UnitOfWork) error {
	// Parse request body
	var pengguna models.InsertUser
	if err := c.BodyParser(&pengguna); err != nil {
		fmt.Printf(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	// Get repository from UOW (no transaction needed for read operation)
	penggunaRepo := uowInstance.PenggunaRepository()

	// Create karyawan
	if err := penggunaRepo.CreatePengguna(&pengguna); err != nil {
		fmt.Printf(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to create pengguna",
			"message": "gagal membuat data pengguna",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Karyawan berhasil dibuat",
	})
}
