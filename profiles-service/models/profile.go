package models

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

// Define Photos as a custom type for JSON array
type Photos []string

// Implement Scan method to decode JSON array into Photos
func (p *Photos) Scan(value interface{}) error {
	if value == nil {
		*p = Photos{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return json.Unmarshal(value.([]byte), p)
	}
	return json.Unmarshal(bytes, p)
}

// Implement Value method to encode Photos into JSON
func (p Photos) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Profile model with Photos as a JSON array
type Profile struct {
	gorm.Model
	UserID uint   `gorm:"not null;uniqueIndex"`
	Age    int    `gorm:"not null"`
	Bio    string `gorm:"type:text"`
	Photos Photos `gorm:"type:jsonb"` // Stored as a JSON array in the database
}
