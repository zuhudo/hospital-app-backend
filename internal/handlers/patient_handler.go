package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hospital-app-backend/internal/models"
	"hospital-app-backend/internal/utils"
)

// In-memory patients store
var patients = []models.Patient{
	{
		ID:          "1",
		UserID:      "1",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@email.com",
		Phone:       "+1234567890",
		DateOfBirth: "1990-05-15",
		Gender:      "Male",
		Address:     "123 Main St, City",
		BloodGroup:  "O+",
		Allergies:   []string{"Penicillin"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          "2",
		UserID:      "2",
		FirstName:   "Jane",
		LastName:    "Smith",
		Email:       "jane.smith@email.com",
		Phone:       "+0987654321",
		DateOfBirth: "1985-08-22",
		Gender:      "Female",
		Address:     "456 Oak Ave, Town",
		BloodGroup:  "A+",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

type PatientHandler struct{}

func NewPatientHandler() *PatientHandler {
	return &PatientHandler{}
}

// GetAll returns all patients
func (h *PatientHandler) GetAll(c *fiber.Ctx) error {
	return utils.SuccessResponse(c, fiber.StatusOK, patients)
}

// GetByID returns a patient by ID
func (h *PatientHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, p := range patients {
		if p.ID == id {
			return utils.SuccessResponse(c, fiber.StatusOK, p)
		}
	}
	return utils.ErrorResponse(c, fiber.StatusNotFound, "Patient not found")
}

// Create adds a new patient
func (h *PatientHandler) Create(c *fiber.Ctx) error {
	var req models.CreatePatientRequest
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	patient := models.Patient{
		ID:          uuid.New().String(),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Phone:       req.Phone,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		Address:     req.Address,
		BloodGroup:  req.BloodGroup,
		InsuranceID: req.InsuranceID,
		Allergies:   req.Allergies,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	patients = append(patients, patient)
	return utils.SuccessResponse(c, fiber.StatusCreated, patient)
}

// Update modifies an existing patient
func (h *PatientHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req models.CreatePatientRequest
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	for i, p := range patients {
		if p.ID == id {
			patients[i].FirstName = req.FirstName
			patients[i].LastName = req.LastName
			patients[i].Email = req.Email
			patients[i].Phone = req.Phone
			patients[i].DateOfBirth = req.DateOfBirth
			patients[i].Gender = req.Gender
			patients[i].Address = req.Address
			patients[i].BloodGroup = req.BloodGroup
			patients[i].InsuranceID = req.InsuranceID
			patients[i].Allergies = req.Allergies
			patients[i].UpdatedAt = time.Now()
			return utils.SuccessResponse(c, fiber.StatusOK, patients[i])
		}
	}
	return utils.ErrorResponse(c, fiber.StatusNotFound, "Patient not found")
}

// Delete removes a patient
func (h *PatientHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, p := range patients {
		if p.ID == id {
			patients = append(patients[:i], patients[i+1:]...)
			return utils.MessageResponse(c, fiber.StatusOK, "Patient deleted successfully")
		}
	}
	return utils.ErrorResponse(c, fiber.StatusNotFound, "Patient not found")
}
