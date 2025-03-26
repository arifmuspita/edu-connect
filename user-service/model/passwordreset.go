package model

import "time"

type PasswordReset struct {
	PasswordResetID uint      `gorm:"primaryKey" json:"password_reset_id"`
	Email           string    `gorm:"type:varchar(100);not null"`
	Token           string    `gorm:"type:varchar(255);not null;unique"`
	ExpiresAt       time.Time `gorm:"not null"`
	Used            bool      `gorm:"default:false"`
	CreatedAt       time.Time
}
