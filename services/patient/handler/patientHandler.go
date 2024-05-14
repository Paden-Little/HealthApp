package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/services/patient/gen"
)

// PatientHandler defines the handlers for the Patient service
type PatientHandler struct {
	db PatientDatabase
}

type PatientDatabase interface {
	Close() error
	CreatePatient(patient *gen.NewPatient) (*gen.Patient, error)
	DeletePatient(id string) error
	GetPatient(id string) (*gen.Patient, error)
	GetPatients() ([]*gen.Patient, error)
}

// NewPatientHandler creates a new PatientHandler
func NewPatientHandler() *PatientHandler {
	return &PatientHandler{}
}

// CheckHealth checks the health of the service
func (h *PatientHandler) CheckHealth(c *gin.Context) {
	c.Status(200)
}

// GetPatient gets a patient by id
func (h *PatientHandler) GetPatient(c *gin.Context, id string) {
	c.JSON(200, nil)
}

// GetPatients gets a list of patients
func (h *PatientHandler) GetPatients(c *gin.Context) {
	c.JSON(200, []int{})
}

// CreatePatient creates a new patient
func (h *PatientHandler) CreatePatient(c *gin.Context) {
	c.JSON(201, nil)
}

// DeletePatient deletes a patient
func (h *PatientHandler) DeletePatient(c *gin.Context, id string) {
	c.Status(204)
}
