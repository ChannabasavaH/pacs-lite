package models

import (
	"time"
)

type Gender string

const (
	GenderMale Gender = "MALE"
	GenderFemale Gender = "FEMALE"
	GenderOther Gender = "OTHER"
)

type Profile struct {
	ID string `gorm:"type:uuid;primaryKey" json:"id"`
	AccountID string `gorm:"type:uuid;uniqueIndex;not null" json:"account_id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName string `gorm:"not null" json:"last_name"`
	Gender Gender `gorm:"type:varchar(10);not null" json:"gender"`
	DOB time.Time `gorm:"type:date;not null" json:"dob"`
	Account *Account
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}