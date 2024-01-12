package model

import (
	"mime/multipart"
	"os"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	regexNumber = `^[0-9]+$`
)

type CallStatusInput struct {
	Word string `json:"word"`
}

func (c *CallStatusInput) ToCallStatus() *CallStatus {
	return &CallStatus{
		Word: c.Word,
	}
}

type ContactInput struct {
	Number      string `json:"number"`
	ContactType string `json:"contact_type"`
}

func (c *ContactInput) ToContact() *Contact {
	return &Contact{
		Number:      c.Number,
		ContactType: c.ContactType,
	}
}

type EmergencyProviderInput struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Name      string `json:"name"`
	EmType    string `json:"em_type"`
}

func (e *EmergencyProviderInput) ToEmergencyProvider() *EmergencyProvider {
	return &EmergencyProvider{
		Longitude: e.Longitude,
		Latitude:  e.Latitude,
		Name:      e.Name,
		EmType:    e.EmType,
	}
}

type EmergencyTypeInput struct {
	Word string `json:"word"`
}

func (e *EmergencyTypeInput) ToEmergencyType() *EmergencyType {
	return &EmergencyType{
		Word: e.Word,
	}
}

type UserRegister struct {
	Name        string `json:"name"`
	NIK         string `json:"nik"`
	PhoneNumber string `json:"phone_number"`
}

func (u UserRegister) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.NIK, validation.Required),
		validation.Field(&u.NIK, validation.Match(regexp.MustCompile(regexNumber))),
		validation.Field(&u.PhoneNumber, validation.Required),
		validation.Field(&u.PhoneNumber, validation.Match(regexp.MustCompile(regexNumber))),
	)
}

func (u *UserRegister) ToUser() *User {
	return &User{
		Name:        u.Name,
		NIK:         u.NIK,
		PhoneNumber: u.PhoneNumber,
	}
}

type UserLogin struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func (u *UserLogin) ToUser() *User {
	return &User{
		PhoneNumber: u.PhoneNumber,
	}
}

type UserUpdate struct {
	Name        string                `form:"name"`
	NIK         string                `form:"nik"`
	PhoneNumber string                `form:"phone_number"`
	Image       *multipart.FileHeader `form:"image"`
}

func (u *UserUpdate) ToUser() *User {
	linkImage := ""
	if u.Image != nil {
		linkImage = os.Getenv("LINK_IMAGE") + u.Image.Filename
	}
	return &User{
		Name:        u.Name,
		NIK:         u.NIK,
		PhoneNumber: u.PhoneNumber,
		Image:       linkImage,
	}
}
