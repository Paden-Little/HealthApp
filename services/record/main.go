package main

import (
	"github.com/services/record/gen"
	"github.com/services/record/handler"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func main() {
	// Create handler
	handler := handler.NewRecordHandler()

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
