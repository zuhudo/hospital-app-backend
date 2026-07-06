package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hospital-app-backend/internal/models"
	"hospital-app-backend/internal/utils"
)

// In-memory appointments store
var appointments = []models.Appointment{
	{
		ID:                   "1",
		PatientID:            "1",
		DoctorID:             "1",
		PatientName:          "John Doe",
		DoctorName:           "Dr. Sarah Johnson",
		DoctorSpecialization: "Cardiology",
		AppointmentDate:      time.Now().Add(24 * time.Hour),
		TimeSlot:             "10:00 AM",
		Status:               "scheduled",
		Type:                 "consultation",
		Reason:               "Annual heart checkup",
		Fee:                  200.00,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	},
}

type AppointmentHandler struct{}

func NewAppointmentHandler() *AppointmentHandler {
	return &AppointmentHandler{}
}

// GetAll returns all appointments
func (h *AppointmentHandler) GetAll(c *fiber.Ctx) error {
	return utils.SuccessResponse(c, fiber.StatusOK, appointments)
}

// GetByID returns an appointment by ID
func (h *AppointmentHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, a := range appointments {
		if a.ID == id {
			return utils.SuccessResponse(c, fiber.StatusOK, a)
		}
	}
	return utils.ErrorResponse(c, fiber.StatusNotFound, "Appointment not found")
}

// Create books a new appointment
func (h *AppointmentHandler) Create(c *fiber.Ctx) error {
	var req models.CreateAppointmentRequest
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	// Parse date
	appointmentDate, err := time.Parse(time.RFC3339, req.AppointmentDate)
	if err != nil {
		appointmentDate, err = time.Parse("2006-01-02", req.AppointmentDate)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid date format")
		}
	}

	// Find doctor name
	doctorName := "Doctor"
	doctorSpec := ""
	for _, d := range doctors {
		if d.ID == req.DoctorID {
			doctorName = "Dr. " + d.FirstName + " " + d.LastName
			doctorSpec = d.Specialization
			break
		}
	}

	// Find patient name
	patientName := "Patient"
	for _, p := range patients {
		if p.ID == req.PatientID {
			patientName = p.FirstName + " " + p.LastName
			break
		}
	}

	appointment := models.Appointment{
		ID:                   uuid.New().String(),
		PatientID:            req.PatientID,
		DoctorID:             req.DoctorID,
		PatientName:          patientName,
		DoctorName:           doctorName,
		DoctorSpecialization: doctorSpec,
		AppointmentDate:      appointmentDate,
		TimeSlot:             req.TimeSlot,
		Status:               "scheduled",
		Type:                 req.Type,
		Reason:               req.Reason,
		Notes:                req.Notes,
		Fee:                  req.Fee,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	appointments = append(appointments, appointment)
	return utils.SuccessResponse(c, fiber.StatusCreated, appointment)
}

// Cancel cancels an appointment
func (h *AppointmentHandler) Cancel(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, a := range appointments {
		if a.ID == id {
			appointments[i].Status = "cancelled"
			appointments[i].UpdatedAt = time.Now()
			return utils.SuccessResponse(c, fiber.StatusOK, appointments[i])
		}
	}
	return utils.ErrorResponse(c, fiber.StatusNotFound, "Appointment not found")
}
