package models

import "time"

type Patient struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	DateOfBirth  string    `json:"date_of_birth"`
	Gender       string    `json:"gender"`
	Address      string    `json:"address,omitempty"`
	BloodGroup   string    `json:"blood_group,omitempty"`
	InsuranceID  string    `json:"insurance_id,omitempty"`
	Allergies    []string  `json:"allergies,omitempty"`
	ProfileImage string    `json:"profile_image,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreatePatientRequest struct {
	FirstName   string   `json:"first_name" validate:"required"`
	LastName    string   `json:"last_name" validate:"required"`
	Email       string   `json:"email" validate:"required,email"`
	Phone       string   `json:"phone" validate:"required"`
	DateOfBirth string   `json:"date_of_birth" validate:"required"`
	Gender      string   `json:"gender" validate:"required"`
	Address     string   `json:"address"`
	BloodGroup  string   `json:"blood_group"`
	InsuranceID string   `json:"insurance_id"`
	Allergies   []string `json:"allergies"`
}
