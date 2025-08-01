package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"crud-test-go/config"
// 	"crud-test-go/models"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// // SetupTestDB initializes a fresh test DB for each test
// func SetupTestDB(t *testing.T) {
// 	dsn := "host=localhost user=postgres password=root dbname=bharat_nxt port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("Failed to connect to test DB: %v", err)
// 	}

// 	// Replace global DB instance
// 	config.DB = db

// 	// // Migrate
// 	// err = config.DB.AutoMigrate(&models.Cohort{})
// 	// if err != nil {
// 	// 	t.Fatalf("Failed to migrate: %v", err)
// 	// }

// 	// Clean table
// 	// config.DB.Exec("DELETE FROM cohort_masters")
// }

// func TestCreateCohort_Success(t *testing.T) {
// 	SetupTestDB(t)

// 	cohort := models.Cohort{
// 		Name:        "Test Cohort",
// 		Priority:    15,
// 		Description: "Test Description",
// 	}
// 	body, _ := json.Marshal(cohort)

// 	req := httptest.NewRequest("POST", "/cohorts", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	CreateCohort(w, req)

// 	res := w.Result()
// 	defer res.Body.Close()

// 	assert.Equal(t, http.StatusCreated, res.StatusCode)

// 	var created models.Cohort
// 	json.NewDecoder(res.Body).Decode(&created)

// 	assert.Equal(t, cohort.Name, created.Name)
// 	assert.NotZero(t, created.ID)
// }

// func TestCreateCohort_InvalidJSON(t *testing.T) {
// 	SetupTestDB(t)

// 	req := httptest.NewRequest("POST", "/cohorts", bytes.NewBuffer([]byte("invalid-json")))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	CreateCohort(w, req)

// 	res := w.Result()
// 	defer res.Body.Close()

// 	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
// }


// func TestGetAllCohorts_WithData(t *testing.T) {
// 	SetupTestDB(t)

// 	// Seed 1 cohort
// 	config.DB.Table("bnxt_user.cohort_master").Create(&models.Cohort{
// 		Name:        "Seed Cohort",
// 		Priority:    16,
// 		Description: "Seeder",
// 	})

// 	req := httptest.NewRequest("GET", "/cohorts", nil)
// 	w := httptest.NewRecorder()

// 	GetAllCohorts(w, req)

// 	res := w.Result()
// 	defer res.Body.Close()

// 	assert.Equal(t, http.StatusOK, res.StatusCode)

// 	var cohorts []models.Cohort
// 	json.NewDecoder(res.Body).Decode(&cohorts)

// 	// assert.Len(t, cohorts, 12) 
// 	assert.Equal(t, "Seed Cohort", cohorts[0].Name)
// }
