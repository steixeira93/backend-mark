package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steixeira93/mark/schemas"
)

func ListAppointmentsHandler(c *gin.Context) {
	appointments := []schemas.Appointment{}
	if err := db.Find(&appointments).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error listing appointments: %s", err.Error()))
		return
	}

	sendSuccess(c, "ListAppointments", appointments)
}
