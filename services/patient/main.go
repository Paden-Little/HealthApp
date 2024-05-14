package main

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"github.com/services/patient/database"

	"github.com/services/patient/gen"
	"github.com/services/patient/handler"
)

func main() {
	// Create database
	db, err := database.NewPatientDatabase()
	if err != nil {
		panic(err)
	}

	// Create handler
	handlers := handler.NewPatientHandler(db)

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
