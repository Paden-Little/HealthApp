package main

import (
	"log"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"

	"github.com/services/provider/database"
	"github.com/services/provider/gen"
	"github.com/services/provider/handler"
)

func main() {
	// Create database
	db, err := database.NewProviderDatabase()
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	defer db.Close()

	// Create handler
	handlers := handler.NewProviderHandler(db)

	// Load open api spec (swagger)
	swagger, err := gen.GetSwagger()
	if err != nil {
		panic(err)
	}

	// Create router
	router := gin.Default()
	router.Use(middleware.OapiRequestValidator(swagger))
	gen.RegisterHandlers(router, handlers)

	// Start server
	router.Run(":3000")
}
