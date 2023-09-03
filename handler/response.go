package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steixeira93/mark/schemas"
)

func sendError(c *gin.Context, code int, msg string) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(c *gin.Context, op string, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type CreateAppointmentResponse struct {
	Message string                      `json:"message"`
	Data    schemas.AppointmentResponse `json:"data"`
}

type DeleteAppointmentResponse struct {
	Message string                      `json:"message"`
	Data    schemas.AppointmentResponse `json:"data"`
}

type ShowAppointmenteResponse struct {
	Message string                      `json:"message"`
	Data    schemas.AppointmentResponse `json:"data"`
}

type ListAppointmentResponse struct {
	Message string                        `json:"message"`
	Data    []schemas.AppointmentResponse `json:"data"`
}

type UpdateAppointmentResponse struct {
	Message string                      `json:"message"`
	Data    schemas.AppointmentResponse `json:"data"`
}
