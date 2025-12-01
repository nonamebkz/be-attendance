package models

import "time"

type Karyawan struct {
	ID           int32      `gorm:"primaryKey;column:id_karyawan;type:serial4" json:"id_karyawan"`
	UserID       *int32     `gorm:"column:user_id" json:"user_id"`
	NIP          *string    `gorm:"column:nip;size:50;unique" json:"nip"`
	Nama         string     `gorm:"column:nama;size:100;not null" json:"nama"`
	TanggalLahir *time.Time `gorm:"column:tanggal_lahir" json:"tanggal_lahir"`
	JenisKelamin *string    `gorm:"column:jenis_kelamin;size:10" json:"jenis_kelamin"`
	JabatanID    *int32     `gorm:"column:jabatan_id" json:"jabatan_id"`
	Alamat       *string    `gorm:"column:alamat" json:"alamat"`
	NoTelepon    *string    `gorm:"column:no_telepon;size:20" json:"no_telepon"`
	Foto         *string    `gorm:"column:foto;size:255" json:"foto"`
}
