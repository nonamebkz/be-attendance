package models

import "time"

type Absensi struct {
	ID          int32      `db:"id_absensi" json:"id_absensi"`
	KaryawanID  int32      `db:"karyawan_id" json:"karyawan_id"`
	Tanggal     time.Time  `db:"tanggal" json:"tanggal"`
	WaktuMasuk  *time.Time `db:"waktu_masuk" json:"waktu_masuk"`
	WaktuPulang *time.Time `db:"waktu_pulang" json:"waktu_pulang"`
	Status      string     `db:"status" json:"status"`
	Keterangan  *string    `db:"keterangan" json:"keterangan"`
	Lokasi      *string    `db:"lokasi" json:"lokasi"`
}
