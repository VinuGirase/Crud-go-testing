package repository

import (
	"crud-test-go/config"
	"crud-test-go/models"
)

func GetAllCohorts() ([]models.Cohort, error) {
	var cohorts []models.Cohort
	db := config.DB
	result := db.Table("bnxt_user.cohort_master").Order("id desc").Find(&cohorts)
	return cohorts, result.Error
}

func GetCohortByID(id uint) (*models.Cohort, error) {
	var cohort models.Cohort
	result := config.DB.Table("bnxt_user.cohort_master").First(&cohort, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cohort, nil
}

func CreateCohort(cohort models.Cohort) (*models.Cohort, error) {
	result := config.DB.Table("bnxt_user.cohort_master").Create(&cohort)
	return &cohort, result.Error
}
