package models

import "time"

type IzinCuti struct {
	ID                int32      `db:"id_izin" json:"id_izin"`
	KaryawanID        int32      `db:"karyawan_id" json:"karyawan_id"`
	TanggalMulai      time.Time  `db:"tanggal_mulai" json:"tanggal_mulai"`
	TanggalSelesai    time.Time  `db:"tanggal_selesai" json:"tanggal_selesai"`
	Jenis             string     `db:"jenis" json:"jenis"`
	Alasan            *string    `db:"alasan" json:"alasan"`
	FileBukti         *string    `db:"file_bukti" json:"file_bukti"`
	StatusPersetujuan string     `db:"status_persetujuan" json:"status_persetujuan"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
	ApprovedAt        *time.Time `db:"approved_at" json:"approved_at"`
}
