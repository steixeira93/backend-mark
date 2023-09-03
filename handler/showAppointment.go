package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steixeira93/mark/schemas"
)

func ShowAppointmentHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	appointment := schemas.Appointment{}
	if err := db.First(&appointment, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("error showing appointment with id: %s", id))
		return
	}

	sendSuccess(c, "ShowAppointment", appointment)
}
