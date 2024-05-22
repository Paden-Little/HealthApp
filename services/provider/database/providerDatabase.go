package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/services/provider/gen"
)

// ProviderDatabase RecordDatabase is the database implementation for the record service
type ProviderDatabase struct {
	db *sqlx.DB
}

// NewProviderDatabase connects to the database using environment variables, and returns a ProviderDatabase.
func NewProviderDatabase() (*ProviderDatabase, error) {
	// Connect to the database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "root"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "provider"),
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

	return &ProviderDatabase{db: db}, nil
}

// Close closes the database connection.
func (d *ProviderDatabase) Close() error {
	return d.db.Close()
}

// GetPassword retrieves a provider's password from the database and returns it.
func (d *ProviderDatabase) GetPassword(email string) (string, error) {
	var password string
	query := "SELECT password FROM provider WHERE email = ?"
	err := d.db.Get(&password, query, email)
	if err != nil {
		return "", fmt.Errorf("failed to get password: %w", err)
	}

	return password, nil
}

// GetProviderID retrieves a provider's ID from the database and returns it.
func (d *ProviderDatabase) GetProviderID(email string) (string, error) {
	var id string
	query := "SELECT id FROM provider WHERE email = ?"
	err := d.db.Get(&id, query, email)
	if err != nil {
		return "", fmt.Errorf("failed to get provider ID: %w", err)
	}

	return id, nil
}

// CreateProvider starts a transaction with the database, inserts a new provider, languages, and services, and commits the transaction. It returns a gen.Provider.
func (d *ProviderDatabase) CreateProvider(provider *gen.NewProvider) (*gen.Provider, error) {
	// Start transaction
	tx, err := d.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	// Insert provider
	id := uuid.New().String()
	query := `INSERT INTO provider (id, name, suffix, bio, email, phone, password) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.Exec(query, id, provider.Name, provider.Suffix, provider.Bio, provider.Email, provider.Phone, provider.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create provider: %w", err)
	}

	// Insert image if provided
	if provider.Image != nil {
		query = `UPDATE provider SET image = ? WHERE id = ?`
		_, err = tx.Exec(query, provider.Image, id)
		if err != nil {
			return nil, fmt.Errorf("failed to update provider image: %w", err)
		}
	}

	// Insert languages
	for _, language := range provider.Languages {
		// First try to get the language ID
		var languageID int
		query = `SELECT id FROM language WHERE language = ?`
		if err := tx.Get(&languageID, query, language); err != nil {
			// If the language does not exist, create it
			if errors.Is(err, sql.ErrNoRows) {
				query = `INSERT INTO language (language) VALUES (?)`
				res, err := tx.Exec(query, language)
				if err != nil {
					return nil, fmt.Errorf("failed to create language: %w", err)
				}
				languageID64, err := res.LastInsertId()
				languageID = int(languageID64)
				if err != nil {
					return nil, fmt.Errorf("failed to get language ID: %w", err)
				}
			} else {
				return nil, fmt.Errorf("failed to get language ID: %w", err)
			}
		}

		// Insert the language into the provider_language table
		query = `INSERT INTO provider_language (provider_id, language_id) VALUES (?, ?)`
		_, err = tx.Exec(query, id, languageID)
		if err != nil {
			return nil, fmt.Errorf("failed to create provider language: %w", err)
		}
	}

	// Insert services
	for _, service := range provider.Services {
		// First try to get the service ID
		var serviceID int
		query = `SELECT id FROM service WHERE service = ?`
		if err := tx.Get(&serviceID, query, service); err != nil {
			// If the service does not exist, create it
			if errors.Is(err, sql.ErrNoRows) {
				query = `INSERT INTO service (service) VALUES (?)`
				res, err := tx.Exec(query, service)
				if err != nil {
					return nil, fmt.Errorf("failed to create service: %w", err)
				}
				serviceID64, err := res.LastInsertId()
				serviceID = int(serviceID64)
				if err != nil {
					return nil, fmt.Errorf("failed to get service ID: %w", err)
				}
			} else {
				return nil, fmt.Errorf("failed to get service ID: %w", err)
			}
		}

		// Insert the service into the provider_service table
		query = `INSERT INTO provider_service (provider_id, service_id) VALUES (?, ?)`
		_, err = tx.Exec(query, id, serviceID)
		if err != nil {
			return nil, fmt.Errorf("failed to create provider service: %w", err)
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Return the created provider
	return &gen.Provider{
		Id:        id,
		Name:      provider.Name,
		Suffix:    provider.Suffix,
		Bio:       provider.Bio,
		Email:     provider.Email,
		Phone:     provider.Phone,
		Services:  provider.Services,
		Languages: provider.Languages,
	}, nil
}

// GetProvider retrieves provider info, services, and languages from the database and returns a gen.Provider.
func (d *ProviderDatabase) GetProvider(id string) (*gen.Provider, error) {
	// Get provider
	var provider gen.Provider
	query := "SELECT * FROM provider WHERE id = ?"
	err := d.db.Get(&provider, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %w", err)
	}

	// Get services and languages
	var services []string
	query = `SELECT s.service FROM provider_service ps 
    	JOIN service s ON ps.service_id = s.id
		WHERE ps.provider_id = ?`
	err = d.db.Select(&services, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider services: %w", err)
	}

	var languages []string
	query = `SELECT l.language FROM provider_language pl 
		JOIN language l ON pl.language_id = l.id 
		WHERE pl.provider_id = ?`
	err = d.db.Select(&languages, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider languages: %w", err)
	}

	// Set services and languages
	provider.Services = services
	provider.Languages = languages

	// Return the provider
	return &provider, nil
}

// GetProviders retrieves providers from the database and returns a slice of gen.Provider.
// Filters by service and name if provided.
func (d *ProviderDatabase) GetProviders(params gen.GetProvidersParams) ([]gen.Provider, error) {
	// Get providers
	var providers []gen.Provider
	var query string
	var queryParams []interface{}
	// If no service is provided, alter query to only filter by name
	if params.Service == nil || len(*params.Service) == 0 {
		query = `SELECT * FROM provider WHERE name LIKE ?`
		queryParams = append(queryParams, fmt.Sprintf("%%%s%%", derefString(params.Name)))
	} else {
		query = `SELECT p.* FROM provider.provider p
		JOIN provider.provider_service ps 
			ON p.id = ps.provider_id 
		JOIN provider.service s 
			ON ps.service_id = s.id 
		WHERE s.service LIKE ? AND p.name LIKE ?`
		queryParams = append(queryParams, fmt.Sprintf("%%%s%%", derefString(params.Service)))
	}

	// Execute query
	err := d.db.Select(&providers, query, queryParams...)
	if err != nil {
		return nil, fmt.Errorf("failed to get providers: %w", err)
	}

	// Get services and languages for each provider
	for i := range providers {
		// Get services and languages
		var services []string
		query = `SELECT s.service FROM provider_service ps
			JOIN service s ON ps.service_id = s.id
			WHERE ps.provider_id = ?`
		err = d.db.Select(&services, query, providers[i].Id)
		if err != nil {
			return nil, fmt.Errorf("failed to get provider services: %w", err)
		}

		fmt.Printf("services: %v\n", services)

		var languages []string
		query = `SELECT l.language FROM provider_language pl
			JOIN language l ON pl.language_id = l.id
			WHERE pl.provider_id = ?`
		err = d.db.Select(&languages, query, providers[i].Id)
		if err != nil {
			return nil, fmt.Errorf("failed to get provider languages: %w", err)
		}

		fmt.Printf("languages: %v\n", languages)

		// Set services and languages
		providers[i].Services = services
		providers[i].Languages = languages
	}

	return providers, nil
}

// DeleteProvider deletes a provider from the database. It returns an error if the operation fails.
func (d *ProviderDatabase) DeleteProvider(id string) error {
	// Delete provider
	query := `DELETE FROM provider WHERE id = ?`
	_, err := d.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete provider: %w", err)
	}

	return nil
}

// UpdateProvider updates a provider in the database. It returns an error if the operation fails.
// Will only update fields that aren't pointers
func (d *ProviderDatabase) UpdateProvider(id string, provider *gen.ProviderUpdate) (*gen.Provider, error) {
	// Start transaction
	tx, err := d.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	// Build query
	query := `UPDATE provider SET `
	var params []any

	if provider.Name != nil {
		query += `name = ?, `
		params = append(params, *provider.Name)
	}
	if provider.Suffix != nil {
		query += `suffix = ?, `
		params = append(params, *provider.Suffix)
	}
	if provider.Bio != nil {
		query += `bio = ?, `
		params = append(params, *provider.Bio)
	}
	if provider.Email != nil {
		query += `email = ?, `
		params = append(params, *provider.Email)
	}
	if provider.Phone != nil {
		query += `phone = ?, `
		params = append(params, *provider.Phone)
	}
	if provider.Password != nil {
		query += `password = ?, `
		params = append(params, *provider.Password)
	}
	if provider.Image != nil {
		query += `image = ?, `
		params = append(params, *provider.Image)
	}
	query = query[:len(query)-2] // Remove trailing comma and space

	// Add WHERE clause
	query += ` WHERE id = ?`
	params = append(params, id)

	// Execute query
	_, err = tx.Exec(query, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to update provider: %w", err)
	}

	// Services
	if provider.Services != nil {
		query = `DELETE FROM provider_service WHERE provider_id = ?`
		_, err = tx.Exec(query, id)
		if err != nil {
			return nil, fmt.Errorf("failed to delete provider services: %w", err)
		}

		for _, service := range *provider.Services {
			// First try to get the service ID
			var serviceID int
			query = `SELECT id FROM service WHERE service = ?`
			if err := tx.Get(&serviceID, query, service); err != nil {
				// If the service does not exist, create it
				if errors.Is(err, sql.ErrNoRows) {
					query = `INSERT INTO service (service) VALUES (?)`
					res, err := tx.Exec(query, service)
					if err != nil {
						return nil, fmt.Errorf("failed to create service: %w", err)
					}
					serviceID64, err := res.LastInsertId()
					serviceID = int(serviceID64)
					if err != nil {
						return nil, fmt.Errorf("failed to get service ID: %w", err)
					}
				} else {
					return nil, fmt.Errorf("failed to get service ID: %w", err)
				}
			}

			// Insert the service into the provider_service table
			query = `INSERT INTO provider_service (provider_id, service_id) VALUES (?, ?)`
			_, err = tx.Exec(query, id, serviceID)
			if err != nil {
				return nil, fmt.Errorf("failed to create provider service: %w", err)
			}
		}
	}

	// Languages
	if provider.Languages != nil {
		query = `DELETE FROM provider_language WHERE provider_id = ?`
		_, err = tx.Exec(query, id)
		if err != nil {
			return nil, fmt.Errorf("failed to delete provider languages: %w", err)
		}

		for _, language := range *provider.Languages {
			// First try to get the language ID
			var languageID int
			query = `SELECT id FROM language WHERE language = ?`
			if err := tx.Get(&languageID, query, language); err != nil {
				// If the language does not exist, create it
				if errors.Is(err, sql.ErrNoRows) {
					query = `INSERT INTO language (language) VALUES (?)`
					res, err := tx.Exec(query, language)
					if err != nil {
						return nil, fmt.Errorf("failed to create language: %w", err)
					}
					languageID64, err := res.LastInsertId()
					languageID = int(languageID64)
					if err != nil {
						return nil, fmt.Errorf("failed to get language ID: %w", err)
					}
				} else {
					return nil, fmt.Errorf("failed to get language ID: %w", err)
				}
			}

			// Insert the language into the provider_language table
			query = `INSERT INTO provider_language (provider_id, language_id) VALUES (?, ?)`
			_, err = tx.Exec(query, id, languageID)
			if err != nil {
				return nil, fmt.Errorf("failed to create provider language: %w", err)
			}
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return d.GetProvider(id)
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func derefString(param *string) string {
	if param == nil {
		return ""
	}
	return *param
}
