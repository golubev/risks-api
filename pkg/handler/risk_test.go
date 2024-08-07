package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	risk "risks-api/pkg/model"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/v1/risks", getRisks)
	r.GET("/v1/risks/:id", getRiskByID)
	r.POST("/v1/risks", postRisk)
	return r
}

func TestGetRisks(t *testing.T) {
	risk.RisksStorage = []risk.Risk{
		{
			ID:       uuid.New(),
			RiskBody: risk.RiskBody{State: risk.Open, Title: "Risk 1", Description: "Some description"},
		},
	}

	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/risks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var risks []risk.Risk
	err := json.Unmarshal(w.Body.Bytes(), &risks)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(risks))
	assert.Equal(t, "Risk 1", risks[0].Title)
}

func TestGetRiskByID(t *testing.T) {
	id := uuid.New()
	risk.RisksStorage = []risk.Risk{
		{
			ID:       id,
			RiskBody: risk.RiskBody{State: risk.Open, Title: "Risk 1", Description: "Some description"},
		},
	}

	router := setupRouter()

	// Test valid ID
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/risks/"+id.String(), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var riskItem risk.Risk
	err := json.Unmarshal(w.Body.Bytes(), &riskItem)
	assert.NoError(t, err)
	assert.Equal(t, "Risk 1", riskItem.Title)

	// Test invalid ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/risks/invalid-id", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPostRisk(t *testing.T) {
	router := setupRouter()

	// Test valid risk
	validRisk := risk.RiskBody{State: risk.Open, Title: "Risk 1", Description: "Some description"}
	jsonValue, _ := json.Marshal(validRisk)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/risks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	var newRisk risk.Risk
	err := json.Unmarshal(w.Body.Bytes(), &newRisk)
	assert.NoError(t, err)
	assert.Equal(t, "Risk 1", newRisk.Title)

	// Test invalid risk - invalid state
	riskInvalidState := risk.RiskBody{State: "Invalid state", Title: "Risk 1", Description: "Some description"}
	jsonValue, _ = json.Marshal(riskInvalidState)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/v1/risks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test invalid risk - missing state
	riskMissingState := risk.RiskBody{State: "", Title: "Risk 1", Description: "Some description"}
	jsonValue, _ = json.Marshal(riskMissingState)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/v1/risks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
