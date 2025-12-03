package models

import "time"

type User struct {
	ID           int32     `db:"id_user" json:"id_user"`
	Username     string    `db:"username" json:"username"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"` // json:"-" menyembunyikan hash saat API response
	NamaLengkap  *string   `db:"nama_lengkap" json:"nama_lengkap"`
	Role         string    `db:"role" json:"role"`
	StatusAkun   string    `db:"status_akun" json:"status_akun"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type InsertUser struct {
	Username     string  `db:"username" json:"username"`
	Email        string  `db:"email" json:"email"`
	PasswordHash string  `db:"password_hash" json:"-"`
	NamaLengkap  *string `db:"nama_lengkap" json:"nama_lengkap"`
	Role         string  `db:"role" json:"role"`
	StatusAkun   string  `db:"status_akun" json:"status_akun"`
	CreatedAt	string  `db:"created_at" json:"created_at"`
	UpdatedAt	string  `db:"updated_at" json:"updated_at"`
}
