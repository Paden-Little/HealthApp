package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/services/patient/gen"
	"github.com/services/patient/security"
)

// PatientHandler defines the handlers for the Patient service
type PatientHandler struct {
	db PatientDatabase
}

// PatientDatabase defines the database operations required for the Patient service
type PatientDatabase interface {
	Close() error
	GetPatientPassword(id string) (string, error)
	GetPatientID(email string) (string, error)
	CreatePatient(patient *gen.NewPatient) (*gen.Patient, error)
	DeletePatient(id string) error
	GetPatient(id string) (*gen.Patient, error)
	GetPatients() ([]*gen.Patient, error)
	UpdatePatient(id string, patient *gen.PatientUpdate) (*gen.Patient, error)
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

// PatientLogin authenticates a patient and returns a JWT. It expects a gen.PatientLogin in the request
func (h *PatientHandler) PatientLogin(c *gin.Context) {
	var patientLogin gen.PatientLogin
	if err := c.BindJSON(&patientLogin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Retrieve patient ID
	patientID, err := h.db.GetPatientID(patientLogin.Email)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	// Retrieve hashed password
	storedPass, err := h.db.GetPatientPassword(patientID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	// Compare hashed password with patient password
	err = security.ComparePasswords(storedPass, patientLogin.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	token, err := security.GenerateJWT(patientID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := gen.JWT{
		Token: token,
		Id:    patientID,
	}

	c.JSON(200, response)
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

	// Hash password
	hashedPass, err := security.HashPassword(newPatient.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	newPatient.Password = hashedPass

	// Create patient
	patient, err := h.db.CreatePatient(&newPatient)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, patient)
}

// DeletePatient calls PatientDatabase.DeletePatient() and returns a 204 status. It expects an id parameter
func (h *PatientHandler) DeletePatient(c *gin.Context, id string) {
	// Auth
	security.AuthMiddleware(c)

	err := h.db.DeletePatient(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}

// UpdatePatient calls PatientDatabase.UpdatePatient() and returns the result. It expects an id parameter and a gen.PatientUpdate in the request body
func (h *PatientHandler) UpdatePatient(c *gin.Context, id string) {
	// Auth
	security.AuthMiddleware(c)

	var patientUpdate gen.PatientUpdate
	if err := c.BindJSON(&patientUpdate); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// hash password if present
	if patientUpdate.Password != nil {
		hashedPass, err := security.HashPassword(*patientUpdate.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		patientUpdate.Password = &hashedPass

	}

	patient, err := h.db.UpdatePatient(id, &patientUpdate)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, patient)
}
