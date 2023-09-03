package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steixeira93/mark/schemas"
)

func UpdateAppointmentHandler(c *gin.Context) {
	request := UpdateAppointmentRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %s", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusNotFound, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	appointment := schemas.Appointment{}

	if err := db.First(&appointment, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "appointment not found")
		return
	}

	// Update appointment
	if request.PatientName != "" {
		appointment.PatientName = request.PatientName
	}
	if request.DoctorName != "" {
		appointment.DoctorName = request.DoctorName
	}
	if request.Date != "" {
		appointment.Date = request.Date
	}
	if request.Confirmed != nil {
		appointment.Confirmed = *request.Confirmed
	}
	if request.Description != "" {
		appointment.Description = request.Description
	}

	if err := db.Save(&appointment).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error updating appointment: %s", err.Error()))
		return
	}

	sendSuccess(c, "UpdateAppointment", appointment)
}
