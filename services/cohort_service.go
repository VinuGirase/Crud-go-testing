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
    Repo repository.CohortRepositoryInterface // ✅ interface, no pointer
}

func NewCohortService(repo repository.CohortRepositoryInterface) *CohortService {
    return &CohortService{Repo: repo}
}


func (s *CohortService) FetchAllCohorts() ([]models.Cohort, error) {
    return s.Repo.GetAllCohorts()
}

func (s *CohortService) FetchCohortByID(id uint) (*models.Cohort, error) {
    return s.Repo.GetCohortByID(id)
}

func (s *CohortService) CreateCohort(cohort models.Cohort) (*models.Cohort, error) {
    return s.Repo.CreateCohort(cohort)
}

