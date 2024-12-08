package models

import "gorm.io/gorm"

type Swipe struct {
	gorm.Model
	SwiperUserID        uint   `gorm:"not null"` // Foreign key to Users table
	SwipedProfileUserID uint   `gorm:"not null"` // Foreign key to Profiles table
	Action              string `gorm:"type:varchar(20);not null"`
}
