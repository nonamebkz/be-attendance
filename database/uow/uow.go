package uow

import (
	"absensi-versevox/repositories"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// UnitOfWork interface untuk mengelola transaksi database
type UnitOfWork interface {
	// Transaction management
	Begin() error
	Commit() error
	Rollback() error
	GetTx() *sqlx.Tx
	IsInTransaction() bool

	// Helper untuk menjalankan fungsi dalam transaksi secara otomatis
	RunInTx(fn func(UnitOfWork) error) error

	// Repository getters - semua repository menggunakan transaksi yang sama
	KaryawanRepository() repositories.KaryawanRepository
	// Tambahkan repository lainnya di sini saat diperlukan
	PenggunaRepository() repositories.PenggunaRepository
}

type unitOfWork struct {
	db           *sqlx.DB
	tx           *sqlx.Tx
	inTx         bool
	karyawanRepo repositories.KaryawanRepository
	penggunaRepo repositories.PenggunaRepository
}

// NewUnitOfWork creates a new UnitOfWork instance
func NewUnitOfWork(db *sqlx.DB) UnitOfWork {
	return &unitOfWork{
		db: db,
	}
}

// Begin starts a new transaction
func (uow *unitOfWork) Begin() error {
	if uow.inTx {
		return errors.New("transaction already started")
	}

	tx, err := uow.db.Beginx()
	if err != nil {
		return err
	}

	uow.tx = tx
	uow.inTx = true

	// Repositories will be lazily initialized in their getters
	return nil
}

// Commit commits the current transaction
func (uow *unitOfWork) Commit() error {
	if !uow.inTx {
		return errors.New("no active transaction to commit")
	}

	err := uow.tx.Commit()
	if err != nil {
		return err
	}

	uow.reset()
	return nil
}

// Rollback rolls back the current transaction
func (uow *unitOfWork) Rollback() error {
	if !uow.inTx {
		return errors.New("no active transaction to rollback")
	}

	err := uow.tx.Rollback()
	if err != nil {
		return err
	}

	uow.reset()
	return nil
}

// RunInTx executes the given function within a transaction
// It handles Begin, Commit, and Rollback (including panic recovery) automatically
func (uow *unitOfWork) RunInTx(fn func(UnitOfWork) error) error {
	if err := uow.Begin(); err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			// Rollback transaction on panic
			if err := uow.Rollback(); err != nil {
				log.Printf("RunInTx: failed to rollback on panic: %v", err)
			}
			// Re-panic to ensure the application knows something went wrong
			panic(p)
		}
	}()

	if err := fn(uow); err != nil {
		// Rollback transaction on error
		if rbErr := uow.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	// Commit transaction on success
	if err := uow.Commit(); err != nil {
		return err
	}

	return nil
}

// reset clears the transaction status and cached repositories
func (uow *unitOfWork) reset() {
	uow.inTx = false
	uow.tx = nil
	uow.karyawanRepo = nil
}

// GetTx returns the current transaction
func (uow *unitOfWork) GetTx() *sqlx.Tx {
	return uow.tx
}

// IsInTransaction checks if there's an active transaction
func (uow *unitOfWork) IsInTransaction() bool {
	return uow.inTx
}

// KaryawanRepository returns the karyawan repository with transaction
func (uow *unitOfWork) KaryawanRepository() repositories.KaryawanRepository {
	// Lazy loading untuk repository
	if uow.karyawanRepo == nil {
		if uow.inTx {
			uow.karyawanRepo = repositories.NewKaryawanRepositoryWithTx(uow.tx)
		} else {
			// Jika tidak dalam transaksi, return repository baru dengan DB biasa
			// Note: Kita tidak menyimpannya di struct karena ini stateless/short-lived
			return repositories.NewKaryawanRepository(uow.db)
		}
	}
	return uow.karyawanRepo
}

// KaryawanRepository returns the karyawan repository with transaction
func (uow *unitOfWork) PenggunaRepository() repositories.PenggunaRepository {
	// Lazy loading untuk repository
	if uow.penggunaRepo == nil {
		if uow.inTx {
			uow.penggunaRepo = repositories.NewPenggunaRepositoryWithTx(uow.tx)
		} else {
			// Jika tidak dalam transaksi, return repository baru dengan DB biasa
			// Note: Kita tidak menyimpannya di struct karena ini stateless/short-lived
			return repositories.NewPenggunaRepository(uow.db)
		}
	}
	return uow.penggunaRepo
}
