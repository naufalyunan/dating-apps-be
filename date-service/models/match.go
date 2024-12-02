package models

import "gorm.io/gorm"

type Match struct {
    gorm.Model
    User1ID   uint   `gorm:"not null"` // First user in the match
    User2ID   uint   `gorm:"not null"` // Second user in the match
}
