package model

import "time"

type Notification struct {
	NotificationID int    `gorm:"primaryKey" json:"notification_id"`
	Email          string `gorm:"not null"`
	Subject        string `gorm:"not null"`
	Message        string `gorm:"not null"`
	Status         string `gorm:"default:'pending'"`
	CreatedAt      time.Time
}
