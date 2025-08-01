package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"crud-test-go/models"
	"github.com/gorilla/mux"
	"crud-test-go/services"
)

func GetAllCohorts(w http.ResponseWriter, r *http.Request) {
	cohorts, err := services.FetchAllCohorts()
	if err != nil {
		http.Error(w, "Failed to fetch cohorts", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cohorts)
}

func GetCohortByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	cohort, err := services.FetchCohortByID(uint(id))
	if err != nil {
		http.Error(w, "Cohort not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cohort)
}

func CreateCohort(w http.ResponseWriter, r *http.Request) {
	var cohort models.Cohort
	if err := json.NewDecoder(r.Body).Decode(&cohort); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	created, err := services.CreateCohort(cohort)
	if err != nil {
		http.Error(w, "Failed to create cohort", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
