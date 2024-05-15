package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"github.com/services/patient/database"
	"log"
	"os"
	"strconv"

	"github.com/services/patient/consulhelper"
	"github.com/services/patient/gen"
	"github.com/services/patient/handler"
)

func main() {
	// Define port
	port, err := strconv.Atoi(getEnv("PORT", "3000"))
	if err != nil {
		log.Fatalf("failed to parse port: %v\n", err)
	}

	// Create database
	db, err := database.NewPatientDatabase()
	if err != nil {
		log.Fatalf("failed to create database: %v\n", err)
	}

	// Create handler
	handlers := handler.NewPatientHandler(db)

	// Load open api spec (swagger)
	swagger, err := gen.GetSwagger()
	if err != nil {
		log.Fatalf("failed to parse swagger spec: %v\n", err)
	}

	// Create router
	router := gin.Default()
	router.Use(middleware.OapiRequestValidator(swagger))
	gen.RegisterHandlers(router, handlers)

	// Register service with consul
	serviceId := consulhelper.RegisterService(
		getEnv("CONSUL_ADDRESS", "localhost:8500"),
		getEnv("CONSUL_SERVICE_NAME", "patient"),
		getEnv("CONSUL_SERVICE_PATH", "/patient"),
		port,
	)
	// Defer deregister service with consul
	defer consulhelper.DeregisterService(
		getEnv("CONSUL_ADDRESS", "localhost:8500"),
		serviceId,
	)

	// Start server
	router.Run(fmt.Sprintf(":%d", port))
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
