package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hospital-app-backend/internal/models"
	"hospital-app-backend/internal/utils"
)

// In-memory records store
var records = []models.MedicalRecord{
	{
		ID:        "1",
		PatientID: "1",
		DoctorID:  "1",
		DoctorName: "Dr. Sarah Johnson",
		VisitDate: time.Now().AddDate(0, -1, 0),
		Diagnosis: "Hypertension Stage 1",
		Symptoms:  "Elevated blood pressure, occasional headaches",
		Treatment: "Lifestyle modifications, medication",
		Prescriptions: []models.Prescription{
			{
				MedicineName: "Lisinopril",
				Dosage:       "10mg",
				Frequency:    "Once daily",
				Duration:     "30 days",
				Instructions: "Take in the morning with water",
			},
		},
		Notes:     "Follow up in 4 weeks",
		CreatedAt: time.Now(),
	},
}

type RecordHandler struct{}

func NewRecordHandler() *RecordHandler {
	return &RecordHandler{}
}

// GetByPatientID returns medical records for a patient
func (h *RecordHandler) GetByPatientID(c *fiber.Ctx) error {
	patientID := c.Params("patientId")
	var patientRecords []models.MedicalRecord
	for _, r := range records {
		if r.PatientID == patientID {
			patientRecords = append(patientRecords, r)
		}
	}
	return utils.SuccessResponse(c, fiber.StatusOK, patientRecords)
}

// Create adds a new medical record
func (h *RecordHandler) Create(c *fiber.Ctx) error {
	var req models.CreateRecordRequest
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	// Find doctor name
	doctorName := "Doctor"
	for _, d := range doctors {
		if d.ID == req.DoctorID {
			doctorName = "Dr. " + d.FirstName + " " + d.LastName
			break
		}
	}

	record := models.MedicalRecord{
		ID:            uuid.New().String(),
		PatientID:     req.PatientID,
		DoctorID:      req.DoctorID,
		DoctorName:    doctorName,
		VisitDate:     time.Now(),
		Diagnosis:     req.Diagnosis,
		Symptoms:      req.Symptoms,
		Treatment:     req.Treatment,
		Prescriptions: req.Prescriptions,
		LabResults:    req.LabResults,
		Notes:         req.Notes,
		CreatedAt:     time.Now(),
	}

	records = append(records, record)
	return utils.SuccessResponse(c, fiber.StatusCreated, record)
}
