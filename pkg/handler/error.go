package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPError struct {
	StatusCode int    `json:"statusCode" example:"500"`
	Message    string `json:"message" example:"Internal server error"`
}

func ErrorHandler(c *gin.Context, recovered interface{}) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, HTTPError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal server error",
	})

}
