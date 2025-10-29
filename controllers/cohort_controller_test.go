package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"crud-test-go/models"
	"crud-test-go/repository"
	"crud-test-go/services"
	"github.com/stretchr/testify/assert"
)

func setupMockController() *CohortController {
	mockRepo := new(repository.MockCohortRepository)
	service := services.NewCohortService(mockRepo)
	return NewCohortController(service)
}

func TestCreateCohort_Success(t *testing.T) {
	controller := setupMockController()
	mockRepo := controller.service.repo.(*repository.MockCohortRepository)

	cohort := models.Cohort{Name: "Test Cohort", Priority: 15, Description: "Test Description"}
	mockRepo.On("CreateCohort", cohort).Return(&cohort, nil)

	body, _ := json.Marshal(cohort)
	req := httptest.NewRequest("POST", "/cohorts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	controller.CreateCohort(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateCohort_InvalidJSON(t *testing.T) {
	controller := setupMockController()

	req := httptest.NewRequest("POST", "/cohorts", bytes.NewBuffer([]byte("invalid-json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	controller.CreateCohort(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetAllCohorts_WithData(t *testing.T) {
	controller := setupMockController()
	mockRepo := controller.service.repo.(*repository.MockCohortRepository)

	mockData := []models.Cohort{
		{Name: "Seed Cohort", Priority: 10, Description: "Mock"},
	}
	mockRepo.On("GetAllCohorts").Return(mockData, nil)

	req := httptest.NewRequest("GET", "/cohorts", nil)
	w := httptest.NewRecorder()

	controller.GetAllCohorts(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var cohorts []models.Cohort
	json.NewDecoder(w.Body).Decode(&cohorts)
	assert.Equal(t, "Seed Cohort", cohorts[0].Name)
	mockRepo.AssertExpectations(t)
}
