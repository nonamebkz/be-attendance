package uow

import (
	"absensi-versevox/models"
	"fmt"
	"log"
)

// ExampleUsage demonstrates how to use UnitOfWork pattern
// This is an example file showing best practices
func ExampleUsage(uow UnitOfWork) error {
	// Mulai transaksi
	if err := uow.Begin(); err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Pastikan rollback jika terjadi error
	defer func() {
		if err := recover(); err != nil {
			if rollbackErr := uow.Rollback(); rollbackErr != nil {
				log.Printf("Error rolling back transaction: %v", rollbackErr)
			}
			panic(err) // re-panic setelah rollback
		}
	}()

	// Contoh: Membuat karyawan baru dalam transaksi
	karyawanRepo := uow.KaryawanRepository()

	newKaryawan := &models.Karyawan{
		Nama: "John Doe",
		// ... field lainnya
	}

	if err := karyawanRepo.CreateKaryawan(newKaryawan); err != nil {
		// Rollback akan dilakukan oleh defer
		return fmt.Errorf("failed to create karyawan: %w", err)
	}

	// Contoh operasi lainnya dalam transaksi yang sama
	// Misalnya: Update absensi, create izin, dll
	// Semua operasi ini akan menggunakan transaksi yang sama

	// Commit semua perubahan
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// ExampleUsageWithErrorHandling demonstrates proper error handling
func ExampleUsageWithErrorHandling(uow UnitOfWork) error {
	if err := uow.Begin(); err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Flag untuk tracking apakah perlu rollback
	shouldRollback := true
	defer func() {
		if shouldRollback {
			if err := uow.Rollback(); err != nil {
				log.Printf("Error rolling back transaction: %v", err)
			}
		}
	}()

	karyawanRepo := uow.KaryawanRepository()

	// Operasi 1
	karyawan := &models.Karyawan{
		Nama: "Jane Doe",
	}
	if err := karyawanRepo.CreateKaryawan(karyawan); err != nil {
		return fmt.Errorf("failed to create karyawan: %w", err)
	}

	// Operasi 2 - jika ini gagal, semua akan di-rollback
	// ... operasi lainnya

	// Jika semua berhasil, commit
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	shouldRollback = false // Jangan rollback karena sudah commit
	return nil
}

// ExampleWithRunInTx demonstrates the recommended way to use UoW
// using the RunInTx helper method
func ExampleWithRunInTx(uow UnitOfWork) error {
	return uow.RunInTx(func(txUow UnitOfWork) error {
		// Semua kode di dalam block ini otomatis berjalan dalam transaksi
		// Jika return error, otomatis Rollback
		// Jika return nil, otomatis Commit
		// Jika panic, otomatis Rollback lalu panic kembali

		repo := txUow.KaryawanRepository()

		karyawan := &models.Karyawan{
			Nama: "Transaction User",
		}

		if err := repo.CreateKaryawan(karyawan); err != nil {
			return err // Akan memicu Rollback
		}

		// Bisa tambahkan operasi repository lain di sini
		// ...

		return nil // Akan memicu Commit
	})
}
