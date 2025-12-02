package handlers

import (
	"absensi-versevox/database/uow"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetKaryawanByID handles GET /karyawan-get endpoint
func GetKaryawanByID(c *fiber.Ctx, uowInstance uow.UnitOfWork) error {
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
	karyawanRepo := uowInstance.KaryawanRepository()

	// Get karyawan by ID
	karyawan, err := karyawanRepo.GetKaryawanByID(int32(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "karyawan not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    karyawan,
	})
}

