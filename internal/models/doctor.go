package models

import "time"

type Doctor struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	Specialization  string    `json:"specialization"`
	Department      string    `json:"department,omitempty"`
	Qualification   string    `json:"qualification,omitempty"`
	ExperienceYears int       `json:"experience_years,omitempty"`
	ConsultationFee float64   `json:"consultation_fee,omitempty"`
	ProfileImage    string    `json:"profile_image,omitempty"`
	Rating          float64   `json:"rating,omitempty"`
	IsAvailable     bool      `json:"is_available"`
	AvailableDays   []string  `json:"available_days,omitempty"`
	Bio             string    `json:"bio,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateDoctorRequest struct {
	FirstName       string   `json:"first_name" validate:"required"`
	LastName        string   `json:"last_name" validate:"required"`
	Email           string   `json:"email" validate:"required,email"`
	Phone           string   `json:"phone" validate:"required"`
	Specialization  string   `json:"specialization" validate:"required"`
	Department      string   `json:"department"`
	Qualification   string   `json:"qualification"`
	ExperienceYears int      `json:"experience_years"`
	ConsultationFee float64  `json:"consultation_fee"`
	AvailableDays   []string `json:"available_days"`
	Bio             string   `json:"bio"`
}
