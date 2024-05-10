package main

import (
	"github.com/services/record/gen"
	"github.com/services/record/handler"
	"github.com/services/record/database"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func main() {
	// Creeate database
	db, err := database.NewRecordDatabase()
	if err != nil {
		panic(err)
	}

	// Create handler
	handler := handler.NewRecordHandler(db)

	// Load oapi spec (swagger)
	swagger, err := gen.GetSwagger()
	if err != nil {
		panic(err)
	}

	// Create router
	router := gin.Default()
	router.Use(middleware.OapiRequestValidator(swagger))
	gen.RegisterHandlers(router, handler)

	// Start server
	router.Run(":3000")
}
