package models

import "time"

type Karyawan struct {
	ID           int32      `db:"id_karyawan" json:"id_karyawan"`
	UserID       *int32     `db:"user_id" json:"user_id"`
	NIP          *string    `db:"nip" json:"nip"`
	Nama         string     `db:"nama" json:"nama"`
	TanggalLahir *time.Time `db:"tanggal_lahir" json:"tanggal_lahir"`
	JenisKelamin *string    `db:"jenis_kelamin" json:"jenis_kelamin"`
	JabatanID    *int32     `db:"jabatan_id" json:"jabatan_id"`
	Alamat       *string    `db:"alamat" json:"alamat"`
	NoTelepon    *string    `db:"no_telepon" json:"no_telepon"`
	Foto         *string    `db:"foto" json:"foto"`
}
