package entities

import (
	"database/sql/driver"
	"encoding/json"
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
	ID     int
	UserID uint
	Age    int
	Bio    string
	Photos Photos
}
