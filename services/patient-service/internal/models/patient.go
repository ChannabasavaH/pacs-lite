package models

import (
	"time"
)

type Gender string

const (
	GenderMale Gender = "MALE"
	GenderFemale Gender = "FEMALE"
	GenderOther Gender = "Other"
)

type Patient struct {
	ID string `gorm:"type:uuid;primaryKey" json:"id"`
	MRN string `gorm:"not null;uniqueIndex:idx_facility_mrn" json:"mrn"`

	FirstName string `gorm:"type:varchar(30);not null" json:"firstName"`
	LastName string `gorm:"type:varchar(30);not null" json:"lastName"`
	Gender Gender `gorm:"not null" json:"gender"`
	DOB time.Time `gorm:"type:date;not null" json:"dob"`

	Phone string `gorm:"uniqueIndex;not null" json:"phone"`
	Email string `gorm:"uniqueIndex;not null" json:"email"`
	Address string `gorm:"type:text;not null" json:"address"`

	CreatedBy string `gorm:"type:uuid;not null" json:"created_by"`
	IsActive bool `gorm:"not null;default:true" json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

