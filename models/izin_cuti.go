package models

import "time"

type IzinCuti struct {
	IDIzin            int32      `gorm:"primaryKey;column:id_izin;type:serial4" json:"id_izin"`
	KaryawanID        int32      `gorm:"column:karyawan_id;not null" json:"karyawan_id"`
	TanggalMulai      time.Time  `gorm:"column:tanggal_mulai;type:date;not null" json:"tanggal_mulai"`
	TanggalSelesai    time.Time  `gorm:"column:tanggal_selesai;type:date;not null" json:"tanggal_selesai"`
	Jenis             string     `gorm:"column:jenis;type:izin_jenis_type;not null" json:"jenis"`
	Alasan            *string    `gorm:"column:alasan" json:"alasan"`
	FileBukti         *string    `gorm:"column:file_bukti;size:255" json:"file_bukti"`
	StatusPersetujuan string     `gorm:"column:status_persetujuan;type:izin_status_type;not null" json:"status_persetujuan"`
	CreatedAt         time.Time  `gorm:"column:created_at;type:timestamptz;not null" json:"created_at"`
	ApprovedAt        *time.Time `gorm:"column:approved_at;type:timestamptz" json:"approved_at"`
}
