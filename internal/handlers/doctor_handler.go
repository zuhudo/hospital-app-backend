package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hospital-app-backend/internal/models"
	"hospital-app-backend/internal/utils"
)

// In-memory doctors store
var doctors = []models.Doctor{
	{
		ID:              "1",
		FirstName:       "Sarah",
		LastName:        "Johnson",
		Email:           "sarah.johnson@hospital.com",
		Phone:           "+1112223333",
		Specialization:  "Cardiology",
		Department:      "Cardiology",
		Qualification:   "MD, FACC",
		ExperienceYears: 15,
		ConsultationFee: 200.00,
		Rating:          4.8,
		IsAvailable:     true,
		AvailableDays:   []string{"Monday", "Wednesday", "Friday"},
		Bio:             "Experienced cardiologist specializing in interventional cardiology.",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	},
	{
		ID:              "2",
		FirstName:       "Michael",
		LastName:        "Chen",
		Email:           "michael.chen@hospital.com",
		Phone:           "+4445556666",
		Specialization:  "Neurology",
		Department:      "Neurology",
		Qualification:   "MD, PhD",
		ExperienceYears: 12,
		ConsultationFee: 180.00,
		Rating:          4.9,
		IsAvailable:     true,
		AvailableDays:   []string{"Tuesday", "Thursday", "Saturday"},
		Bio:             "Board-certified neurologist with expertise in movement disorders.",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	},
	{
		ID:              "3",
		FirstName:       "Emily",
		LastName:        "Williams",
		Email:           "emily.williams@hospital.com",
		Phone:           "+7778889999",
		Specialization:  "Pediatrics",
		Department:      "Pediatrics",
		Qualification:   "MD, FAAP",
		ExperienceYears: 8,
		ConsultationFee: 150.00,
		Rating:          4.7,
		IsAvailable:     true,
		AvailableDays:   []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"},
		Bio:             "Compassionate pediatrician dedicated to child health and wellness.",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	},
}

type DoctorHandler struct{}

func NewDoctorHandler() *DoctorHandler {
	return &DoctorHandler{}
}

// GetAll returns all doctors
func (h *DoctorHandler) GetAll(c *fiber.Ctx) error {
	return utils.SuccessResponse(c, fiber.StatusOK, doctors)
}

// GetByID returns a doctor by ID
func (h *DoctorHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, d := range doctors {
		if d.ID == id {
			return utils.SuccessResponse(c, fiber.StatusOK, d)
		}
	}
	return utils.ErrorResponse(c, fiber.StatusNotFound, "Doctor not found")
}

// Create adds a new doctor
func (h *DoctorHandler) Create(c *fiber.Ctx) error {
	var req models.CreateDoctorRequest
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	doctor := models.Doctor{
		ID:              uuid.New().String(),
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		Email:           req.Email,
		Phone:           req.Phone,
		Specialization:  req.Specialization,
		Department:      req.Department,
		Qualification:   req.Qualification,
		ExperienceYears: req.ExperienceYears,
		ConsultationFee: req.ConsultationFee,
		AvailableDays:   req.AvailableDays,
		Bio:             req.Bio,
		IsAvailable:     true,
		Rating:          0,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	doctors = append(doctors, doctor)
	return utils.SuccessResponse(c, fiber.StatusCreated, doctor)
}
