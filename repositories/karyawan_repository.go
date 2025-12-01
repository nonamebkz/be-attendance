package repositories

import (
	"absensi-versevox/models"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

// DBTx interface untuk fleksibilitas menggunakan *sqlx.DB atau *sqlx.Tx
type DBTx interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type KaryawanRepository interface {
	CreateKaryawan(karyawan *models.Karyawan) error
	GetKaryawanByID(id int32) (*models.Karyawan, error)
	UpdateKaryawan(karyawan *models.Karyawan) error
	DeleteKaryawan(id int32) error
}

type karyawanRepository struct {
	db DBTx
}

// NewKaryawanRepository creates a new repository with *sqlx.DB
func NewKaryawanRepository(db *sqlx.DB) KaryawanRepository {
	return &karyawanRepository{db: db}
}

// NewKaryawanRepositoryWithTx creates a new repository with *sqlx.Tx for transactions
func NewKaryawanRepositoryWithTx(tx *sqlx.Tx) KaryawanRepository {
	return &karyawanRepository{db: tx}
}

func (r *karyawanRepository) CreateKaryawan(karyawan *models.Karyawan) error {
	query := `INSERT INTO karyawan (user_id, nip, nama, tanggal_lahir, jenis_kelamin, jabatan_id, alamat, no_telepon, foto) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id_karyawan`
	row := r.db.QueryRow(query, karyawan.UserID, karyawan.NIP, karyawan.Nama, karyawan.TanggalLahir, karyawan.JenisKelamin, karyawan.JabatanID, karyawan.Alamat, karyawan.NoTelepon, karyawan.Foto)
	err := row.Scan(&karyawan.ID)
	if err != nil {
		log.Printf("Error creating karyawan: %v", err)
		return err
	}
	return nil
}

func (r *karyawanRepository) GetKaryawanByID(id int32) (*models.Karyawan, error) {
	var karyawan models.Karyawan
	query := `SELECT * FROM karyawan WHERE id_karyawan = $1`
	err := r.db.Get(&karyawan, query, id)
	if err != nil {
		log.Printf("Error getting karyawan by ID: %v", err)
		return nil, err
	}
	return &karyawan, nil
}

func (r *karyawanRepository) UpdateKaryawan(karyawan *models.Karyawan) error {
	query := `UPDATE karyawan SET user_id=$1, nip=$2, nama=$3, tanggal_lahir=$4, jenis_kelamin=$5, jabatan_id=$6, alamat=$7, no_telepon=$8, foto=$9 WHERE id_karyawan=$10`
	_, err := r.db.Exec(query, karyawan.UserID, karyawan.NIP, karyawan.Nama, karyawan.TanggalLahir, karyawan.JenisKelamin, karyawan.JabatanID, karyawan.Alamat, karyawan.NoTelepon, karyawan.Foto, karyawan.ID)
	if err != nil {
		log.Printf("Error updating karyawan: %v", err)
		return err
	}
	return nil
}

func (r *karyawanRepository) DeleteKaryawan(id int32) error {
	query := `DELETE FROM karyawan WHERE id_karyawan = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting karyawan: %v", err)
		return err
	}
	return nil
}
