package main

import "github.com/gin-gonic/gin"

type Server struct{}

func (s *Server) GetAppointmentAppointmentId(c *gin.Context, AppointmentId string) {
	c.JSON(200, "Blah blah blah blah")
}

func (s *Server) DeleteAppointmentAppointmentId(c *gin.Context, AppointmentId string) {

}

func (s *Server) PatchAppointmentAppointmentId(c *gin.Context, AppointmentId string) {

}

func (s *Server) PostAppointment(c *gin.Context) {

}

func NewServer() *Server {
	return &Server{}
}
