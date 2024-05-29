package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"

	dal "appointment/DAL"
	gen "appointment/gen"
)

type Server struct {
	db                 dal.AppointmentDatabase
	postResponse       gen.AppointmentResponse
	badRequestResponse gen.BadRequestResponse
	notFoundResponse   gen.NotFoundResponse
}

func (s *Server) GetAppointmentAppointmentId(c *gin.Context, AppointmentId string) {
	appointment, err := s.db.SelectAppointmentsByProviderOrPatient(AppointmentId)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			c.JSON(http.StatusNotFound, s.notFoundResponse)
		default:
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (s *Server) DeleteAppointmentAppointmentId(c *gin.Context, AppointmentId string) {
	if err := s.db.DeleteAppointmentById(AppointmentId); err != nil {
		switch {
		case err == sql.ErrNoRows:
			c.JSON(http.StatusNotFound, s.notFoundResponse)
		default:
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, "")
}

func (s *Server) PatchAppointmentAppointmentId(c *gin.Context, AppointmentId string) {
	var AppUpdate gen.Appointment
	if err := c.ShouldBindJSON(&AppUpdate); err != nil {
		*s.badRequestResponse.Message = fmt.Sprintf(*s.badRequestResponse.Message, err.Error())
		c.JSON(http.StatusBadRequest, s.badRequestResponse)
		return
	}

	AppPreUpdate, err := s.db.SelectAppointmentById(AppointmentId)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			c.JSON(http.StatusNotFound, s.notFoundResponse)
		default:
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	if AppUpdate.Date == nil {
		AppUpdate.Date = AppPreUpdate.Date
	}
	if AppUpdate.StartTime == nil {
		AppUpdate.StartTime = AppPreUpdate.StartTime
	}
	if AppUpdate.EndTime == nil {
		AppUpdate.EndTime = AppPreUpdate.EndTime
	}
	if AppUpdate.Provider == nil {
		AppUpdate.Provider = AppPreUpdate.Provider
	}
	if AppUpdate.Service == nil {
		AppUpdate.Service = AppPreUpdate.Service
	}
	if AppUpdate.Description == nil {
		AppUpdate.Description = AppPreUpdate.Description
	}

	if err := s.db.UpdateAppointmentById(AppointmentId, &AppUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func (s *Server) PostAppointment(c *gin.Context) {
	var App gen.Appointment
	if err := c.ShouldBindJSON(&App); err != nil {
		*s.badRequestResponse.Message = fmt.Sprintf(*s.badRequestResponse.Message, err.Error())
		c.JSON(http.StatusBadRequest, s.badRequestResponse)
	}

	if app, err := s.db.InsertAppointment(&App); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		App = *app
	}

	s.postResponse.Id = App.Id
	c.JSON(http.StatusCreated, s.postResponse)
}

func (s *Server) CheckHealth(c *gin.Context) {
	c.Status(200)
}

func NewServer(db dal.AppointmentDatabase) *Server {
	BadRequestMessage := "Could not proccess your request. One or more incorrect fields: "
	AppointmentCreationMessage := "Successfully Scheduled Appointment"
	NotFoundMessage := "Could not find appointment with ID: "

	return &Server{
		db: db,
		badRequestResponse: gen.BadRequestResponse{
			Message: &BadRequestMessage,
		},
		postResponse: gen.AppointmentResponse{
			Message: &AppointmentCreationMessage,
		},
		notFoundResponse: gen.NotFoundResponse{
			Message: &NotFoundMessage,
		},
	}
}
