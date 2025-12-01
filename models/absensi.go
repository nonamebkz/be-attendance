package models

import "time"

type Absensi struct {
	ID          int32      `gorm:"primaryKey;column:id_absensi;type:serial4" json:"id_absensi"`
	KaryawanID  int32      `gorm:"column:karyawan_id;not null" json:"karyawan_id"`
	Tanggal     time.Time  `gorm:"column:tanggal;type:date;not null" json:"tanggal"`
	WaktuMasuk  *time.Time `gorm:"column:waktu_masuk;type:timestamptz" json:"waktu_masuk"`
	WaktuPulang *time.Time `gorm:"column:waktu_pulang;type:timestamptz" json:"waktu_pulang"`
	Status      string     `gorm:"column:status;type:absensi_status_type;not null" json:"status"`
	Keterangan  *string    `gorm:"column:keterangan" json:"keterangan"`
	Lokasi      *string    `gorm:"column:lokasi;size:255" json:"lokasi"`
}
