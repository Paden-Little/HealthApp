package handler

import "github.com/gin-gonic/gin"

type RecordHandler struct{
	db Database
}

type Database interface {
	// CreateConditionRecord()
	// DeleteConditionRecord()
	// GetConditionRecord()
	// UpdateConditionRecord()
	// CreateEncounterRecord()
	// DeleteEncounterRecord()
	// GetEncounterRecord()
	// UpdateEncounterRecord()
	// CreateImmunizationRecord()
	// DeleteImmunizationRecord()
	// GetImmunizationRecord()
	// UpdateImmunizationRecord()
	// CreateMedicationRecord()
	// DeleteMedicationRecord()
	// GetMedicationRecord()
	// UpdateMedicationRecord()
	// CreatePatientRecord()
	// DeletePatientRecord()
	// GetPatientRecord()
	// UpdatePatientRecord()
	// CreateProviderRecord()
	// DeleteProviderRecord()
	// GetProviderRecord()
	// UpdateProviderRecord()
}

func NewRecordHandler(db Database) *RecordHandler {
	return &RecordHandler{db: db}
}

func (h *RecordHandler) CheckHealth(c *gin.Context) {
    c.Status(200)
}

func (h *RecordHandler) CreateConditionRecord(c *gin.Context) {
    c.Status(200)
}

func (h *RecordHandler) DeleteConditionRecord(c *gin.Context, conditionId string) {
    c.Status(200)
}

func (h *RecordHandler) GetConditionRecord(c *gin.Context, conditionId string) {
    c.Status(200)
}

func (h *RecordHandler) UpdateConditionRecord(c *gin.Context, conditionId string) {
    c.Status(200)
}

func (h *RecordHandler) CreateEncounterRecord(c *gin.Context) {
    c.Status(200)
}

func (h *RecordHandler) DeleteEncounterRecord(c *gin.Context, encounterId string) {
    c.Status(200)
}

func (h *RecordHandler) GetEncounterRecord(c *gin.Context, encounterId string) {
    c.Status(200)
}

func (h *RecordHandler) UpdateEncounterRecord(c *gin.Context, encounterId string) {
    c.Status(200)
}

func (h *RecordHandler) CreateImmunizationRecord(c *gin.Context) {
    c.Status(200)
}

func (h *RecordHandler) DeleteImmunizationRecord(c *gin.Context, immunizationId string) {
    c.Status(200)
}

func (h *RecordHandler) GetImmunizationRecord(c *gin.Context, immunizationId string) {
    c.Status(200)
}

func (h *RecordHandler) UpdateImmunizationRecord(c *gin.Context, immunizationId string) {
    c.Status(200)
}

func (h *RecordHandler) CreateMedicationRecord(c *gin.Context) {
    c.Status(200)
}

func (h *RecordHandler) DeleteMedicationRecord(c *gin.Context, medicationId string) {
    c.Status(200)
}

func (h *RecordHandler) GetMedicationRecord(c *gin.Context, medicationId string) {
    c.Status(200)
}

func (h *RecordHandler) UpdateMedicationRecord(c *gin.Context, medicationId string) {
    c.Status(200)
}

func (h *RecordHandler) CreatePatientRecord(c *gin.Context) {
    c.Status(200)
}

func (h *RecordHandler) DeletePatientRecord(c *gin.Context, patientId string) {
    c.Status(200)
}

func (h *RecordHandler) GetPatientRecord(c *gin.Context, patientId string) {
    c.Status(200)
}

func (h *RecordHandler) UpdatePatientRecord(c *gin.Context, patientId string) {
    c.Status(200)
}

func (h *RecordHandler) CreateProviderRecord(c *gin.Context) {
    c.Status(200)
}

func (h *RecordHandler) DeleteProviderRecord(c *gin.Context, providerId string) {
    c.Status(200)
}

func (h *RecordHandler) GetProviderRecord(c *gin.Context, providerId string) {
    c.Status(200)
}

func (h *RecordHandler) UpdateProviderRecord(c *gin.Context, providerId string) {
    c.Status(200)
}