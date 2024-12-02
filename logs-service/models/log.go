package models

import "gorm.io/gorm"

type ActivityLog struct {
    gorm.Model
    UserID        uint   `gorm:"not null"` // Foreign key to Users table
    ActionType    string `gorm:"type:enum('swipe', 'purchase');not null"`
    ActionDetails string `gorm:"type:text"`
}
