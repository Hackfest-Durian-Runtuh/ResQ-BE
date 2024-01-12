package model

import (
	"time"

	"gorm.io/gorm"
)

// CallStatusModel represents the structure for the CallStatus model.
type CallStatus struct {
	ID        int            `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Word      string         `json:"word"`
}

// Contact represents the structure for the Contact .
type Contact struct {
	ID                  int            `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"-"`
	EmergencyProviderID int            `json:"emergency_provider_id"`
	Number              string         `json:"number"`
	ContactType         string         `json:"contact_type"`
}

// EmergencyProvider represents the structure for the EmergencyProvider .
type EmergencyProvider struct {
	ID        int            `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Longitude string         `json:"longitude"`
	Latitude  string         `json:"latitude"`
	Name      string         `json:"name"`
	EmType    string         `json:"em_type"`
}

// EmergencyType represents the structure for the EmergencyType .
type EmergencyType struct {
	ID        int            `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Word      string         `json:"word"`
}

// User represents the structure for the User .
type User struct {
	ID          int            `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `json:"name"`
	NIK         string         `json:"nik"`
	PhoneNumber string         `json:"phone_number"`
	IsAdmin     bool           `json:"is_admin" gorm:"default:false"`
	Image       string         `json:"image"`
}
