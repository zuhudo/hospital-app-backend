package models

import "time"

type Appointment struct {
	ID                   string    `json:"id"`
	PatientID            string    `json:"patient_id"`
	DoctorID             string    `json:"doctor_id"`
	PatientName          string    `json:"patient_name,omitempty"`
	DoctorName           string    `json:"doctor_name,omitempty"`
	DoctorSpecialization string    `json:"doctor_specialization,omitempty"`
	AppointmentDate      time.Time `json:"appointment_date"`
	TimeSlot             string    `json:"time_slot"`
	Status               string    `json:"status"` // scheduled, confirmed, in_progress, completed, cancelled, no_show
	Type                 string    `json:"type"`   // consultation, follow_up, emergency, lab_test, surgery
	Reason               string    `json:"reason,omitempty"`
	Notes                string    `json:"notes,omitempty"`
	Fee                  float64   `json:"fee,omitempty"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type CreateAppointmentRequest struct {
	PatientID       string  `json:"patient_id" validate:"required"`
	DoctorID        string  `json:"doctor_id" validate:"required"`
	AppointmentDate string  `json:"appointment_date" validate:"required"`
	TimeSlot        string  `json:"time_slot" validate:"required"`
	Type            string  `json:"type" validate:"required"`
	Reason          string  `json:"reason"`
	Notes           string  `json:"notes"`
	Fee             float64 `json:"fee"`
}
