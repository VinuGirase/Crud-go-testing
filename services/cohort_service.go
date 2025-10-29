// package services

// import (
// 	"crud-test-go/models"
// 	"crud-test-go/repository"
// )

// func FetchAllCohorts() ([]models.Cohort, error) {
// 	return repository.GetAllCohorts()
// }

// func FetchCohortByID(id uint) (*models.Cohort, error) {
// 	return repository.GetCohortByID(id)
// }

// func CreateCohort(cohort models.Cohort) (*models.Cohort, error) {
// 	return repository.CreateCohort(cohort)
// }

package services

import (
	"crud-test-go/models"
	"crud-test-go/repository"
)

type CohortService struct {
	repo *repository.CohortRepository
}

func NewCohortService(repo *repository.CohortRepository) *CohortService {
	return &CohortService{repo: repo}
}

func (s *CohortService) FetchAllCohorts() ([]models.Cohort, error) {
	return s.repo.GetAllCohorts()
}

func (s *CohortService) FetchCohortByID(id uint) (*models.Cohort, error) {
	return s.repo.GetCohortByID(id)
}

func (s *CohortService) CreateCohort(cohort models.Cohort) (*models.Cohort, error) {
	return s.repo.CreateCohort(cohort)
}
