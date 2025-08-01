// package routes

// import (
// 	// "net/http"
// 	"github.com/gorilla/mux"
// 	"crud-test-go/controllers"
// )

// func SetupRoutes() *mux.Router {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
// 	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")

// 	return r
// }

package routes

import (
	"github.com/gorilla/mux"
	"crud-test-go/controllers"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// User APIs
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/cleanup", controllers.CleanupTestUsers).Methods("DELETE")


	// Cohort APIs
	r.HandleFunc("/cohorts", controllers.GetAllCohorts).Methods("GET")
	r.HandleFunc("/cohorts/{id}", controllers.GetCohortByID).Methods("GET")
	r.HandleFunc("/cohorts", controllers.CreateCohort).Methods("POST")

	return r
}
