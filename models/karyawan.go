package models

type Karyawan struct {
	ID           int32   `db:"id_karyawan" json:"id_karyawan"`
	UserID       *int32  `db:"user_id" json:"user_id"`
	NIP          *string `db:"nip" json:"nip"`
	Nama         string  `db:"nama" json:"nama"`
	TanggalLahir *string `db:"tanggal_lahir" json:"tanggal_lahir"`
	JenisKelamin *string `db:"jenis_kelamin" json:"jenis_kelamin"`
	JabatanID    *int32  `db:"jabatan_id" json:"jabatan_id"`
	Alamat       *string `db:"alamat" json:"alamat"`
	NoTelepon    *string `db:"no_telepon" json:"no_telepon"`
	Foto         *string `db:"foto" json:"foto"`
}

type InsertKaryawan struct {
	UserID       *int32  `db:"user_id" json:"user_id"`
	NIP          *string `db:"nip" json:"nip"`
	Nama         string  `db:"nama" json:"nama"`
	TanggalLahir *string `db:"tanggal_lahir" json:"tanggal_lahir"`
	JenisKelamin *string `db:"jenis_kelamin" json:"jenis_kelamin"`
	JabatanID    *int32  `db:"jabatan_id" json:"jabatan_id"`
	Alamat       *string `db:"alamat" json:"alamat"`
	NoTelepon    *string `db:"no_telepon" json:"no_telepon"`
	Foto         *string `db:"foto" json:"foto"`
}

type InsertKaryawanRequest struct {
	NIP          *string `json:"nip"`
	Nama         string  `json:"nama"`
	TanggalLahir *string `json:"tanggal_lahir"`
	JenisKelamin *string `json:"jenis_kelamin"`
	JabatanID    *int32  `json:"jabatan_id"`
	Alamat       *string `json:"alamat"`
	NoTelepon    *string `json:"no_telepon"`
	Foto         *string `json:"foto"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	PasswordHash string  `json:"password"`
	Role         string  `json:"role"`
	StatusAkun   string  `json:"status_akun"`
}
