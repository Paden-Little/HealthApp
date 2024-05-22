package main

import (
	"log"
	"net/http"

	gen "appointment/gen"

	"github.com/gin-gonic/gin"
)

func main() {
	server := NewServer()

	r := gin.Default()

	gen.RegisterHandlers(r, server)

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Print("it is up")

	log.Fatal(s.ListenAndServe())
}
