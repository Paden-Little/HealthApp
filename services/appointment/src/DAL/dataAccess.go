package DAL

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	gen "appointment/gen"
)

type AppointmentDatabase struct {
	db *sqlx.DB
}

// NewPatientDatabase connects to the database using environment variables, and returns a PatientDatabase.
func NewAppointmentDatabase() (*AppointmentDatabase, error) {
	// Connect to MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "root"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "appointment"),
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

	return &AppointmentDatabase{db: db}, nil
}

func (db *AppointmentDatabase) UpdateAppointmentById(UUID string, App *gen.Appointment) error {
	tx, err := db.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	appointmentUpdate := `UPDATE appointment SET date = ?, start_time = ?, end_time = ?, provider = ?, patient = ?, service = ?, description = ? WHERE id = ?`

	_, err = tx.Exec(appointmentUpdate, App.Date, App.StartTime, App.EndTime, App.Provider, App.Patient, App.Service, App.Description, App.Id)
	if err != nil {
		return fmt.Errorf("failed to execute UPDATE statement: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (db *AppointmentDatabase) DeleteAppointmentById(UUID string) error {
	tx, err := db.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	appointmentDelete := `DELETE FROM appointment WHERE id =?`

	_, err = tx.Exec(appointmentDelete, UUID)
	if err != nil {
		return fmt.Errorf("failed to execute DELETE statement: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (db *AppointmentDatabase) SelectAppointmentsByProviderOrPatient(UUID string) ([]*gen.Appointment, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	var appoints []*gen.Appointment

	// Adjusted SQL query to select appointments where the provider or patient matches the given UUID
	appointmentSelect := `SELECT id, date, start_time, end_time, provider, patient, service, description FROM appointment WHERE provider =? OR patient =?`

	rows, err := tx.Query(appointmentSelect, UUID, UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SELECT statement: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var appoint gen.Appointment
		if err := rows.Scan(&appoint.Id, &appoint.Date, &appoint.StartTime, &appoint.EndTime, &appoint.Provider, &appoint.Patient, &appoint.Service, &appoint.Description); err != nil {
			return nil, fmt.Errorf("failed to scan result: %w", err)
		}
		appoints = append(appoints, &appoint)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return appoints, nil
}

func (db *AppointmentDatabase) SelectAppointmentById(UUID string) (*gen.Appointment, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	var appoint gen.Appointment

	appointmentSelect := `SELECT id, date, start_time, end_time, provider, patient, service, description FROM appointment WHERE id =?`

	rows, err := tx.Query(appointmentSelect, UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SELECT statement: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	if err := rows.Scan(&appoint.Id, &appoint.Date, &appoint.StartTime, &appoint.EndTime, &appoint.Provider, &appoint.Patient, &appoint.Service, &appoint.Description); err != nil {
		return nil, fmt.Errorf("failed to scan result: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &appoint, nil
}

func (db *AppointmentDatabase) InsertAppointment(app *gen.Appointment) (*gen.Appointment, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var UUID string
	UUID = uuid.New().String()

	appointmentInsert := `INSERT INTO appointment (id, date, start_time, end_time, provider, patient, service, description) VALUES (?,?,?,?,?,?,?,?)`
	_, err = tx.Exec(appointmentInsert, UUID, app.Date, app.StartTime, app.EndTime, app.Provider, app.Patient, app.Service, app.Description)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	app.Id = &UUID

	return app, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
