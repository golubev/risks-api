package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "risks-api/docs"
)

func (s State) IsValid() bool {
	switch s {
	case Open, Closed, Accepted, Investigating:
		return true
	}

	return false
}

type State string

const (
	Open          State = "open"
	Closed        State = "closed"
	Accepted      State = "accepted"
	Investigating State = "investigating"
)

// swagger:model
type RiskBody struct {
	State       State  `json:"state" example:"investigating" binding:"required" extensions:"x-order=1"`
	Title       string `json:"title" example:"CVE-2022-29217" extensions:"x-order=2"`
	Description string `json:"description" example:"python-jose through 3.3.0 has algorithm confusion with OpenSSH ECDSA keys and other key formats." extensions:"x-order=3"`
}

// swagger:model
type Risk struct {
	ID uuid.UUID `json:"id" example:"add736b0-516b-401c-a4ee-bfa00812bb52" extensions:"x-order=0"`
	RiskBody
}

type HTTPError struct {
	StatusCode int    `json:"statusCode" example:"500"`
	Message    string `json:"message" example:"Internal server error"`
}

var risks = []Risk{}

// @Summary Get Risks
// @Description get Risks
// @Accept json
// @Produce json
// @Success 200 {array} Risk
// @Failure 500 {object} HTTPError
// @Router /v1/risks [get]
func getRisks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, risks)
}

// @Summary Get Risk By ID
// @Description get Risk by ID
// @Accept json
// @Produce json
// @Param id path string true "Risk ID"
// @Success 200 {object} Risk
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /v1/risks/{id} [get]
func getRiskByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range risks {
		if a.ID.String() == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "risk not found"})
}

// @Summary Post Risk
// @Description post Risk
// @Accept json
// @Produce json
// @Param input body RiskBody true "risk data"
// @Success 201 {object} Risk
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /v1/risks [post]
func postRisk(c *gin.Context) {
	var newRisk Risk

	if err := c.BindJSON(&newRisk); err != nil || !newRisk.State.IsValid() {
		c.AbortWithStatusJSON(http.StatusBadRequest, HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request",
		})
		return
	}

	newRisk.ID = uuid.New()

	risks = append(risks, newRisk)
	c.IndentedJSON(http.StatusCreated, newRisk)
}

func ErrorHandler(c *gin.Context, recovered interface{}) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, HTTPError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal server error",
	})

}

func main() {
	router := gin.Default()

	router.Use(gin.CustomRecovery(ErrorHandler))

	v1 := router.Group("/v1")
	{
		v1.GET("/risks", getRisks)
		v1.GET("/risks/:id", getRiskByID)
		v1.POST("/risks", postRisk)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8080")

	docs.SwaggerInfo.BasePath = "/"
}
