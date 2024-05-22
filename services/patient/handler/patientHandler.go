package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/services/patient/gen"
)

// PatientHandler defines the handlers for the Patient service
type PatientHandler struct {
	db PatientDatabase
}

// PatientDatabase defines the database operations required for the Patient service
type PatientDatabase interface {
	Close() error
	CreatePatient(patient *gen.NewPatient) (*gen.Patient, error)
	DeletePatient(id string) error
	GetPatient(id string) (*gen.Patient, error)
	GetPatients() ([]*gen.Patient, error)
}

// NewPatientHandler creates a new PatientHandler. It requires a PatientDatabase
func NewPatientHandler(db PatientDatabase) *PatientHandler {
	return &PatientHandler{
		db: db,
	}
}

// CheckHealth returns a 200 status, indicating the service is healthy
func (h *PatientHandler) CheckHealth(c *gin.Context) {
	c.Status(200)
}

// GetPatient calls PatientDatabase.GetPatient() and returns the result. It expects an id parameter
func (h *PatientHandler) GetPatient(c *gin.Context, id string) {
	patient, err := h.db.GetPatient(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, patient)
}

// GetPatients calls PatientDatabase.GetPatients() and returns the result. It expects no parameters
func (h *PatientHandler) GetPatients(c *gin.Context) {
	patients, err := h.db.GetPatients()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, patients)
}

// CreatePatient calls PatientDatabase.CreatePatient() and returns the result. It expects a gen.NewPatient in the request body
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

// DeletePatient calls PatientDatabase.DeletePatient() and returns a 204 status. It expects an id parameter
func (h *PatientHandler) DeletePatient(c *gin.Context, id string) {
	err := h.db.DeletePatient(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}
