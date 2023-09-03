package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	intializeRoutes(router)

	router.Run()
}