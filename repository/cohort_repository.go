// package repository

// import (
// 	"crud-test-go/config"
// 	"crud-test-go/models"
// )

// func GetAllCohorts() ([]models.Cohort, error) {
// 	var cohorts []models.Cohort
// 	db := config.DB
// 	result := db.Table("bnxt_user.cohort_master").Order("id desc").Find(&cohorts)
// 	return cohorts, result.Error
// }

// func GetCohortByID(id uint) (*models.Cohort, error) {
// 	var cohort models.Cohort
// 	result := config.DB.Table("bnxt_user.cohort_master").First(&cohort, id)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &cohort, nil
// }

// func CreateCohort(cohort models.Cohort) (*models.Cohort, error) {
// 	result := config.DB.Table("bnxt_user.cohort_master").Create(&cohort)
// 	return &cohort, result.Error
// }

package repository

import (
	"crud-test-go/models"
	"gorm.io/gorm"
)

// Wrap repository with struct
type CohortRepository struct {
	db *gorm.DB
}

type CohortRepositoryInterface interface {
    GetAllCohorts() ([]models.Cohort, error)
    GetCohortByID(id uint) (*models.Cohort, error)
    CreateCohort(cohort models.Cohort) (*models.Cohort, error)
}

func NewCohortRepository(db *gorm.DB) *CohortRepository {
	return &CohortRepository{db: db}
}

func (r *CohortRepository) GetAllCohorts() ([]models.Cohort, error) {
	var cohorts []models.Cohort
	result := r.db.Table("bnxt_user.cohort_master").Order("id desc").Find(&cohorts)
	return cohorts, result.Error
}

func (r *CohortRepository) GetCohortByID(id uint) (*models.Cohort, error) {
	var cohort models.Cohort
	result := r.db.Table("bnxt_user.cohort_master").First(&cohort, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cohort, nil
}

func (r *CohortRepository) CreateCohort(cohort models.Cohort) (*models.Cohort, error) {
	result := r.db.Table("bnxt_user.cohort_master").Create(&cohort)
	return &cohort, result.Error
}
