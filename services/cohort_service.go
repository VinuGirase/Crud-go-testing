package services

import (
	"crud-test-go/models"
	"crud-test-go/repository"
)

func FetchAllCohorts() ([]models.Cohort, error) {
	return repository.GetAllCohorts()
}

func FetchCohortByID(id uint) (*models.Cohort, error) {
	return repository.GetCohortByID(id)
}

func CreateCohort(cohort models.Cohort) (*models.Cohort, error) {
	return repository.CreateCohort(cohort)
}
