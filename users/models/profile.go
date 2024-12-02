package models

import (
	"gorm.io/gorm"
)

type Profile struct {
    gorm.Model
    UserID uint   `gorm:"not null;uniqueIndex"`
    Age    int    `gorm:"not null"`
    Bio    string `gorm:"type:text"`
    Photos string `gorm:"type:jsonb"` // Storing URLs as JSON array
}

