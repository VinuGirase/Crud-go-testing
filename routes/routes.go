// // package routes

// // import (
// // 	// "net/http"
// // 	"github.com/gorilla/mux"
// // 	"crud-test-go/controllers"
// // )

// // func SetupRoutes() *mux.Router {
// // 	r := mux.NewRouter()

// // 	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
// // 	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")

// // 	return r
// // }

// package routes

// import (
// 	"github.com/gorilla/mux"
// 	"crud-test-go/controllers"
// )

// func SetupRoutes() *mux.Router {
// 	r := mux.NewRouter()

// 	// User APIs
// 	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
// 	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
// 	r.HandleFunc("/users/cleanup", controllers.CleanupTestUsers).Methods("DELETE")

// 	// Cohort APIs
// 	r.HandleFunc("/cohorts", controllers.GetAllCohorts).Methods("GET")
// 	r.HandleFunc("/cohorts/{id}", controllers.GetCohortByID).Methods("GET")
// 	r.HandleFunc("/cohorts", controllers.CreateCohort).Methods("POST")

//		return r
//	}



package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"crud-test-go/config"
	"crud-test-go/controllers"
	"crud-test-go/repository"
	"crud-test-go/services"
)

// SetupRoutes sets up all API routes with injected controllers
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Initialize repository, service, and controller with dependency injection
	cohortRepo := repository.NewCohortRepository(config.DB)      // pass *gorm.DB
	cohortService := services.NewCohortService(cohortRepo)       // pass repo to service
	cohortController := controllers.NewCohortController(cohortService) // pass service to controller

	// User APIs
	r.HandleFunc("/users", controllers.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users", controllers.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/cleanup", controllers.CleanupTestUsers).Methods(http.MethodDelete)

	// Cohort APIs
	cohortRouter := r.PathPrefix("/cohorts").Subrouter()
	cohortRouter.HandleFunc("", cohortController.GetAllCohorts).Methods(http.MethodGet)
	cohortRouter.HandleFunc("/{id:[0-9]+}", cohortController.GetCohortByID).Methods(http.MethodGet)
	cohortRouter.HandleFunc("", cohortController.CreateCohort).Methods(http.MethodPost)

	return r
}
