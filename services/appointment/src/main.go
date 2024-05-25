package main

import (
	"log"
	"net/http"
	"os"

	gen "appointment/gen"

	dal "appointment/DAL"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := dal.NewAppointmentDatabase()
	if err != nil {
		log.Fatalf("failed to create database: %v\n", err)
	}

	server := NewServer(*db)

	r := gin.Default()

	gen.RegisterHandlers(r, server)

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Print("it is up")

	// Register service with consul
	serviceId := RegisterService(
		getEnv("CONSUL_ADDRESS", "localhost:8500"),
		getEnv("CONSUL_SERVICE_NAME", "appointment"),
		getEnv("CONSUL_SERVICE_PATH", "/appointment"),
		8080,
	)
	// Defer deregister service with consul
	defer DeregisterService(
		getEnv("CONSUL_ADDRESS", "localhost:8500"),
		serviceId,
	)

	log.Fatal(s.ListenAndServe())
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
