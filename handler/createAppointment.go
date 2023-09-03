package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steixeira93/mark/schemas"
)

func CreateAppointmentHandler(c *gin.Context) {
	request := CreateAppointmentRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %s", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	appointment := schemas.Appointment{
		DoctorName:  request.DoctorName,
		PatientName: request.PatientName,
		Date:        request.Date,
		Confirmed:   *request.Confirmed,
		Description: request.Description,
	}

	if err := db.Create(&appointment).Error; err != nil {
		logger.Errorf("Error creating appointment: %s", err.Error())
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "CreateAppointment", appointment)

}
