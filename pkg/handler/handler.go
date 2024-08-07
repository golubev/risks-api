package handler

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.CustomRecovery(ErrorHandler))

	v1 := router.Group("/v1")
	{
		v1.GET("/risks", getRisks)
		v1.GET("/risks/:id", getRiskByID)
		v1.POST("/risks", postRisk)
	}

	return router
}
