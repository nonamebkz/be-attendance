package handlers

import (
	"absensi-versevox/database/uow"
	"absensi-versevox/models"
	"fmt"
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

func CreateKaryawanByID(c *fiber.Ctx, uowInstance uow.UnitOfWork) error {
	// Parse request body
	var karyawan models.InsertKaryawan
	if err := c.BodyParser(&karyawan); err != nil {
		fmt.Printf(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	// Get repository from UOW (no transaction needed for read operation)
	karyawanRepo := uowInstance.KaryawanRepository()

	// Create karyawan
	if err := karyawanRepo.CreateKaryawan(&karyawan); err != nil {
		fmt.Printf(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to create karyawan",
			"message": "gagal membuat data karyawan",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Karyawan berhasil dibuat",
	})
}

func UpdateKaryawanByID(c *fiber.Ctx, uowInstance uow.UnitOfWork) error {
	//get id from query parameter
	idParam := c.Query("id")
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id parameter is required",
		})
	}

	// Convert string to int32
	_, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id format",
		})
	}

	// Parse request body
	var karyawan models.Karyawan
	if err := c.BodyParser(&karyawan); err != nil {
		fmt.Printf(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// Get repository from UOW (no transaction needed for read operation)
	karyawanRepo := uowInstance.KaryawanRepository()

	// Update karyawan
	if err := karyawanRepo.UpdateKaryawan(&karyawan); err != nil {
		fmt.Printf(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to update karyawan",
			"message": "gagal memperbarui data karyawan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Karyawan berhasil diperbarui",
		"data":    karyawan,
	})
}

func DeleteKaryawanByID(c *fiber.Ctx, uowInstance uow.UnitOfWork) error {
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

	// Delete karyawan
	if err := karyawanRepo.DeleteKaryawan(int32(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete karyawan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Karyawan berhasil dihapus",
	})
}
