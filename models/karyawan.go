package models

import "time"

type Karyawan struct {
	UserID       uint      `gorm:"column:user_id;unique;not null" json:"user_id"`
	NIP          string    `gorm:"column:nip;unique;not null" json:"nip"`
	Nama         string    `gorm:"column:nama;not null" json:"nama"`
	TanggalLahir time.Time `gorm:"column:tanggal_lahir" json:"tanggal_lahir"`
	JenisKelamin string    `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	JabatanID    uint      `gorm:"column:jabatan_id" json:"jabatan_id"`
	Alamat       string    `gorm:"column:alamat" json:"alamat"`
	NoTelepon    string    `gorm:"column:no_telepon" json:"no_telepon"`
	Foto         string    `gorm:"column:foto" json:"foto"`
}
