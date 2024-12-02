package models

import "gorm.io/gorm"

type ActivityLog struct {
	gorm.Model
	UserID        uint   `gorm:"not null"` // Foreign key to Users table
	ActionType    string `gorm:"type:varchar(20);not null"`
	ActionDetails string `gorm:"type:text"`
}
