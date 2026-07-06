package models

import "time"

type MedicalRecord struct {
	ID           string         `json:"id"`
	PatientID    string         `json:"patient_id"`
	DoctorID     string         `json:"doctor_id"`
	DoctorName   string         `json:"doctor_name,omitempty"`
	VisitDate    time.Time      `json:"visit_date"`
	Diagnosis    string         `json:"diagnosis"`
	Symptoms     string         `json:"symptoms,omitempty"`
	Treatment    string         `json:"treatment,omitempty"`
	Prescriptions []Prescription `json:"prescriptions,omitempty"`
	LabResults   []string       `json:"lab_results,omitempty"`
	Notes        string         `json:"notes,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
}

type Prescription struct {
	MedicineName string `json:"medicine_name"`
	Dosage       string `json:"dosage"`
	Frequency    string `json:"frequency"`
	Duration     string `json:"duration"`
	Instructions string `json:"instructions,omitempty"`
}

type CreateRecordRequest struct {
	PatientID     string         `json:"patient_id" validate:"required"`
	DoctorID      string         `json:"doctor_id" validate:"required"`
	Diagnosis     string         `json:"diagnosis" validate:"required"`
	Symptoms      string         `json:"symptoms"`
	Treatment     string         `json:"treatment"`
	Prescriptions []Prescription `json:"prescriptions"`
	LabResults    []string       `json:"lab_results"`
	Notes         string         `json:"notes"`
}
