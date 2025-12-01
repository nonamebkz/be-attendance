package models

import "time"

type Reset struct {
	ID   int32      `gorm:"primaryKey;column:id_reset;type:serial4" json:"id_reset"`
	UserID    int32      `gorm:"column:user_id;not null" json:"user_id"`
	Token     string     `gorm:"column:token;size:255;not null" json:"token"`
	ExpiredAt time.Time  `gorm:"column:expired_at;type:timestamptz;not null" json:"expired_at"`
	UsedAt    *time.Time `gorm:"column:used_at;type:timestamptz" json:"used_at"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamptz;not null" json:"created_at"`
}
