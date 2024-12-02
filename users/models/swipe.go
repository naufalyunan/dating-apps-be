package models

import "gorm.io/gorm"

type Swipe struct {
    gorm.Model
    SwiperID        uint   `gorm:"not null"` // Foreign key to Users table
    SwipedProfileID uint   `gorm:"not null"` // Foreign key to Profiles table
    Action          string `gorm:"type:enum('like', 'pass');not null"`
}