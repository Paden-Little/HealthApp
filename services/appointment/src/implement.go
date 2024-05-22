package main

import (
	"github.com/gin-gonic/gin"

	dal "appointment/DAL"
)

type Server struct {
	db dal.AppointmentDatabase
}

func (s *Server) GetAppointmentAppointmentId(c *gin.Context, AppointmentId string) {
	c.JSON(200, "Blah blah blah blah")
}

func (s *Server) DeleteAppointmentAppointmentId(c *gin.Context, AppointmentId string) {

}

func (s *Server) PatchAppointmentAppointmentId(c *gin.Context, AppointmentId string) {

}

func (s *Server) PostAppointment(c *gin.Context) {

}

func (s *Server) CheckHealth(c *gin.Context) {
	c.Status(200)
}

func NewServer(db dal.AppointmentDatabase) *Server {
	return &Server{
		db: db,
	}
}
