package models

import "time"

type Reset struct {
	ID        int32      `db:"id_reset" json:"id_reset"`
	UserID    int32      `db:"user_id" json:"user_id"`
	Token     string     `db:"token" json:"token"`
	ExpiredAt time.Time  `db:"expired_at" json:"expired_at"`
	UsedAt    *time.Time `db:"used_at" json:"used_at"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
}
