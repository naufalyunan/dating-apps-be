package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"size:100;not null"`
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	IsPremium    bool   `gorm:"default:false"`
	IsVerified   bool   `gorm:"default:false"`
	// Profiles     Profile `gorm:"foreignKey:UserID"`
	// Swipes       []Swipe `gorm:"foreignKey:SwiperID"`
	// Matches      []Match `gorm:"foreignKey:User1ID"`
	// ActivityLogs []ActivityLog  `gorm:"foreignKey:UserID"`
}
