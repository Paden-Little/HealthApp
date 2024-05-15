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
func NewPatientHandler(db PatientDatabase) *PatientHandler {
	return &PatientHandler{
		db: db,
	}
}

// CheckHealth checks the health of the service
func (h *PatientHandler) CheckHealth(c *gin.Context) {
	c.Status(200)
}

// GetPatient gets a patient by id
func (h *PatientHandler) GetPatient(c *gin.Context, id string) {
	patient, err := h.db.GetPatient(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, patient)
}

// GetPatients gets a list of patients
func (h *PatientHandler) GetPatients(c *gin.Context) {
	patients, err := h.db.GetPatients()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, patients)
}

// CreatePatient creates a new patient
func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var newPatient gen.NewPatient
	if err := c.BindJSON(&newPatient); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	patient, err := h.db.CreatePatient(&newPatient)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, patient)
}

// DeletePatient deletes a patient
func (h *PatientHandler) DeletePatient(c *gin.Context, id string) {
	err := h.db.DeletePatient(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}
