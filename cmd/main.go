package main

import (
	docs "risks-api/docs"
	handler "risks-api/pkg/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Risks API
// @version 1.0
// @description Risks REST API

func main() {
	router := handler.SetUpRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:8080")

	docs.SwaggerInfo.BasePath = "/"
}
