package models

import "time"

type AppUser struct {
	IDUser       int32     `gorm:"primaryKey;column:id_user;type:serial4" json:"id_user"`
	Username     string    `gorm:"column:username;size:50;not null;unique" json:"username"`
	Email        string    `gorm:"column:email;size:100;not null;unique" json:"email"`
	PasswordHash string    `gorm:"column:password_hash;size:255;not null" json:"-"` // json:"-" menyembunyikan hash saat API response
	NamaLengkap  *string   `gorm:"column:nama_lengkap;size:100" json:"nama_lengkap"`
	Role         string    `gorm:"column:role;type:role_type;not null" json:"role"`
	StatusAkun   string    `gorm:"column:status_akun;type:akun_status_type;not null" json:"status_akun"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamptz;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamptz;not null" json:"updated_at"`
}
