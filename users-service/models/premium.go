package models

import "gorm.io/gorm"

type PremiumFeature struct {
    gorm.Model
    UserID      uint   `gorm:"not null"` // Foreign key to Users table
    FeatureType string `gorm:"type:enum('no_swipe_limit', 'verified_badge');not null"`
    ExpiresAt   string `gorm:"not null"`
}
