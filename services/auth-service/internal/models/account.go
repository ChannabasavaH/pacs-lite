package models

import (
	"time"
)

type AccountType string

const (
	AccountTypeRadiologist AccountType = "RADIOLOGIST"
	AccountTypeHospitalITAdmin AccountType = "HOSPITAL_IT_ADMIN"
	AccountTypeNetworkAdmin AccountType = "NETWORK_ADMIN"
)

type Account struct {
	ID string `gorm:"type:uuid;primaryKey;" json:"id" `
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:" not null" json:"-"`
	AccountType AccountType `gorm:"type:varchar(30);not null" json:"account_type"`
	Profile *Profile `gorm:"constraint:OnDelete:CASCADE" json:"profile"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
