package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	risk "risks-api/pkg/model"
)

// @Summary Get Risks
// @Description get Risks
// @Accept json
// @Produce json
// @Success 200 {array} risk.Risk
// @Failure 500 {object} HTTPError
// @Router /v1/risks [get]
func getRisks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, risk.RisksStorage)
}

// @Summary Get Risk By ID
// @Description Get Risk By ID
// @Accept json
// @Produce json
// @Param id path string true "Risk ID"
// @Success 200 {object} risk.Risk
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /v1/risks/{id} [get]
func getRiskByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range risk.RisksStorage {
		if a.ID.String() == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, HTTPError{
		StatusCode: http.StatusNotFound,
		Message:    "Risk Not Found",
	})
}

// @Summary Create Risk
// @Description Create Risk
// @Accept json
// @Produce json
// @Param input body risk.RiskBody true "Risk Data"
// @Success 201 {object} risk.Risk
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /v1/risks [post]
func postRisk(c *gin.Context) {
	var newRisk risk.Risk

	if err := c.BindJSON(&newRisk); err != nil || !newRisk.State.IsValid() {
		c.AbortWithStatusJSON(http.StatusBadRequest, HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request",
		})
		return
	}

	newRisk.ID = uuid.New()

	risk.RisksStorage = append(risk.RisksStorage, newRisk)
	c.IndentedJSON(http.StatusCreated, newRisk)
}
