package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/services/patient/gen"
)

// ProviderDatabase defines the database for the Provider service
type PatientDatabase struct {
	db *sqlx.DB
}

// NewProviderDatabase creates a new ProviderDatabase
func NewProviderDatabase() (*PatientDatabase, error) {
	// Connect to the database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "root"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "patient"),
	)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Ping to verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PatientDatabase{db: db}, nil
}

// Close closes the database connection
func (d *PatientDatabase) Close() error {
	return d.db.Close()
}

// CreatePatient creates a new patient
func (d *PatientDatabase) CreatePatient(patient *gen.NewPatient) (*gen.Patient, error) {
	// Start transaction
	tx, err := d.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	// Check for language
	if derefString(patient.Language) == "" {
		language := "english"
		patient.Language = &language
	}

	// Try to get languageId
	var languageId int
	query := `SELECT l.language FROM provider.language l WHERE l.language = ?`
	if err := tx.Get(&languageId, query, *patient.Language); err != nil {
		// If we can't find the language, insert it
		if errors.Is(err, sql.ErrNoRows) {
			query = `INSERT INTO provider.language (language) VALUES (?)`
			res, err := tx.Exec(query, *patient.Language)
			if err != nil {
				return nil, fmt.Errorf("failed to insert language: %w", err)
			}
			languageId64, err := res.LastInsertId()
			languageId = int(languageId64)
			if err != nil {
				return nil, fmt.Errorf("failed to get language ID: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to get language: %w", err)
		}
	}

	// Insert patient
	id := uuid.New().String()
	query = `INSERT INTO patient.patient (id, firstname, lastname, email, phone, language, birth, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.Exec(query, id, patient.Firstname, patient.Lastname, patient.Email, patient.Phone, languageId, patient.Birth, patient.Gender)
	if err != nil {
		return nil, fmt.Errorf("failed to insert patient: %w", err)
	}

}

// getEnv gets an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
