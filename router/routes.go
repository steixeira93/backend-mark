package router

import (
	"github.com/gin-gonic/gin"
	"github.com/steixeira93/mark/handler"
)

func intializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.GET("/appointments", handler.ListAppointmentsHandler)
		v1.POST("/appointment", handler.CreateAppointmentHandler)
		v1.DELETE("/appointment", handler.DeleteAppointmentHandler)
		v1.PUT("/appointment", handler.UpdateAppointmentHandler)
		v1.GET("/appointment/:id", handler.ShowAppointmentHandler)
	}
}
