package repository

import (
	"crud-test-go/models"
	"github.com/stretchr/testify/mock"
)

type MockCohortRepository struct {
	mock.Mock
}

func (m *MockCohortRepository) GetAllCohorts() ([]models.Cohort, error) {
	args := m.Called()
	return args.Get(0).([]models.Cohort), args.Error(1)
}

func (m *MockCohortRepository) GetCohortByID(id uint) (*models.Cohort, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Cohort), args.Error(1)
}

func (m *MockCohortRepository) CreateCohort(cohort models.Cohort) (*models.Cohort, error) {
	args := m.Called(cohort)
	return args.Get(0).(*models.Cohort), args.Error(1)
}
