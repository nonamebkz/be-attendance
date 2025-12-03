package repositories

import (
	"absensi-versevox/models"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

// DBTx interface untuk fleksibilitas menggunakan *sqlx.DB atau *sqlx.Tx
type DbTx interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type PenggunaRepository interface {
	CreatePengguna(pengguna *models.InsertUser) error
	GetPenggunaByID(id int32) (*models.User, error)
	UpdatePengguna(pengguna *models.User) error
	DeletePengguna(id int32) error
}

type penggunaRepository struct {
	db DBTx
}

// NewKaryawanRepository creates a new repository with *sqlx.DB
func NewPenggunaRepository(db *sqlx.DB) PenggunaRepository {
	return &penggunaRepository{db: db}
}

// NewKaryawanRepositoryWithTx creates a new repository with *sqlx.Tx for transactions
func NewPenggunaRepositoryWithTx(tx *sqlx.Tx) PenggunaRepository {
	return &penggunaRepository{db: tx}
}

func (r *penggunaRepository) CreatePengguna(pengguna *models.InsertUser) error {
	query := `INSERT INTO app_user (username, email, password_hash, nama_lengkap, role,status_akun) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, pengguna.Username, pengguna.Email, pengguna.PasswordHash, pengguna.NamaLengkap, pengguna.Role, pengguna.StatusAkun)
	if err != nil {
		return err
	}
	return nil
}

func (r *penggunaRepository) GetPenggunaByID(id int32) (*models.User, error) {
	var pengguna models.User
	query := `SELECT * FROM app_user WHERE id_user = $1`
	err := r.db.Get(&pengguna, query, id)
	if err != nil {
		log.Printf("Error getting pengguna by ID: %v", err)
		return nil, err
	}
	return &pengguna, nil
}

func (r *penggunaRepository) UpdatePengguna(pengguna *models.User) error {
	query := `UPDATE karyawan SET username=$1, email=$2, password_hash=$3, nama_lengkap=$4, role=$5, status_akun=$6 WHERE id_user=$7`
	_, err := r.db.Exec(query, pengguna.Username, pengguna.Email, pengguna.PasswordHash, pengguna.NamaLengkap, pengguna.Role, pengguna.StatusAkun, pengguna.ID)
	if err != nil {
		log.Printf("Error updating pengguna: %v", err)
		return err
	}
	return nil
}

func (r *penggunaRepository) DeletePengguna(id int32) error {
	query := `DELETE FROM app_user WHERE id_user = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting pengguna: %v", err)
		return err
	}
	return nil
}
