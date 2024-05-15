package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"os"

	"github.com/services/patient/gen"
)

// PatientDatabase implements the database operations required for the Patient service
type PatientDatabase struct {
	db *sqlx.DB
}

// NewPatientDatabase connects to the database using environment variables, and returns a PatientDatabase.
func NewPatientDatabase() (*PatientDatabase, error) {
	// Connect to MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "root"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "patient"),
	)
	fmt.Println("Connecting to database with DSN:", dsn)
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

// Close closes the database connection.
func (d *PatientDatabase) Close() error {
	return d.db.Close()
}

// CreatePatient starts a transaction with the database, inserts a new patient, allergies, and prescriptions, and commits the transaction.
// If the language does not already exist, a new one is created. It returns a gen.Patient.
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
	query := `SELECT l.id FROM provider.language l WHERE l.language = ?`
	if err := tx.Get(&languageId, query, *patient.Language); err != nil {
		// If we can't find the language, insert it
		if errors.Is(err, sql.ErrNoRows) {
			query = `INSERT INTO provider.language (language) VALUES (?)`
			res, err := tx.Exec(query, *patient.Language)
			if err != nil {
				return nil, fmt.Errorf("failed to insert language: %w", err)
			}
			// mfw I have to convert int64 to int (╯°□°)╯︵ ┻━┻
			languageId64, err := res.LastInsertId()
			if err != nil {
				return nil, fmt.Errorf("failed to get language ID: %w", err)
			}
			languageId = int(languageId64)
		} else {
			return nil, fmt.Errorf("failed to get language: %w", err)
		}
	}

	// Insert into patient table
	id := uuid.New().String()
	query = `INSERT INTO patient.patient (id, firstname, lastname, email, phone, language, birth, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.Exec(query, id, patient.Firstname, patient.Lastname, patient.Email, patient.Phone, languageId, patient.Birth, patient.Gender)
	if err != nil {
		return nil, fmt.Errorf("failed to insert patient: %w", err)
	}

	// Handle nil allergies
	allergies := *patient.Allergies
	if allergies == nil {
		allergies = []gen.Allergy{}
	}
	// Insert allergies into allergy table
	for _, allergy := range allergies {
		query = `INSERT INTO patient.allergy(patient_id, name, description) VALUES (?, ?, ?)`
		_, err = tx.Exec(query, id, allergy.Name, allergy.Description)
	}

	// Handle nil prescriptions
	prescriptions := *patient.Prescriptions
	if prescriptions == nil {
		prescriptions = []gen.Prescription{}
	}
	// Insert prescriptions into prescription table
	for _, prescription := range prescriptions {
		query = `INSERT INTO patient.prescription(provider_id, patient_id, name, dosage, frequency, start, end) VALUES (?, ?, ?, ?, ?, ?, ?)`
		_, err = tx.Exec(query, prescription.ProviderId, id, prescription.Name, prescription.Dosage, prescription.Frequency, prescription.Start, prescription.End)
		if err != nil {
			return nil, fmt.Errorf("failed to insert prescription: %w", err)
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &gen.Patient{
		Id:            id,
		Firstname:     patient.Firstname,
		Lastname:      patient.Lastname,
		Email:         patient.Email,
		Phone:         patient.Phone,
		Language:      patient.Language,
		Gender:        patient.Gender,
		Birth:         patient.Birth,
		Allergies:     &allergies,
		Prescriptions: &prescriptions,
	}, nil
}

// GetPatient retrieves patient data, allergies, and prescriptions from the database. It returns a gen.Patient.
func (d *PatientDatabase) GetPatient(id string) (*gen.Patient, error) {
	// Get patient
	var patient gen.Patient
	query := `SELECT p.id, p.firstname, p.lastname, p.email, p.phone, l.language
		FROM patient.patient p
		JOIN provider.language l ON p.language = l.id
		WHERE p.id = ?`
	if err := d.db.Get(&patient, query, id); err != nil {
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}

	// Get allergies
	var allergies []gen.Allergy
	query = `SELECT a.name, a.description FROM patient.allergy a WHERE a.patient_id = ?`
	if err := d.db.Select(&allergies, query, id); err != nil {
		return nil, fmt.Errorf("failed to get allergies: %w", err)
	}

	// Get prescriptions
	var prescriptions []gen.Prescription
	query = `SELECT pr.provider_id, pr.name, pr.dosage, pr.frequency, pr.start, pr.end
		FROM patient.prescription pr WHERE pr.patient_id = ?`
	if err := d.db.Select(&prescriptions, query, id); err != nil {
		return nil, fmt.Errorf("failed to get prescriptions: %w", err)
	}

	// Set allergies and prescriptions
	patient.Allergies = &allergies
	patient.Prescriptions = &prescriptions

	return &patient, nil
}

// GetPatients retrieves all patients, allergies, and prescriptions from the database. It returns a slice of gen.Patient.
func (d *PatientDatabase) GetPatients() ([]*gen.Patient, error) {
	var patients []*gen.Patient
	query := `SELECT p.id, p.firstname, p.lastname, p.email, p.phone, l.language, p.gender, p.birth
		FROM patient.patient p
		JOIN provider.language l ON p.language = l.id`
	if err := d.db.Select(&patients, query); err != nil {
		return nil, fmt.Errorf("failed to get patients: %w", err)
	}

	// Get allergies and prescriptions
	for i := range patients {
		// Get allergies
		var allergies []gen.Allergy
		query = `SELECT a.name, a.description FROM patient.allergy a WHERE a.patient_id = ?`
		if err := d.db.Select(&allergies, query, patients[i].Id); err != nil {
			return nil, fmt.Errorf("failed to get allergies: %w", err)
		}
		patients[i].Allergies = &allergies

		// Get prescriptions
		var prescriptions []gen.Prescription
		query = `SELECT pr.provider_id, pr.name, pr.dosage, pr.frequency, pr.start, pr.end
			FROM patient.prescription pr WHERE pr.patient_id = ?`
		if err := d.db.Select(&prescriptions, query, patients[i].Id); err != nil {
			return nil, fmt.Errorf("failed to get prescriptions: %w", err)
		}
		patients[i].Prescriptions = &prescriptions
	}

	return patients, nil
}

// DeletePatient deletes a patient from the database. It returns an error if the operation fails.
func (d *PatientDatabase) DeletePatient(id string) error {
	// Delete patient
	query := `DELETE FROM patient.patient WHERE id = ?`
	if _, err := d.db.Exec(query, id); err != nil {
		return fmt.Errorf("failed to delete patient: %w", err)
	}

	return nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// derefString dereferences a string pointer, returning an empty string if the pointer is nil
func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
