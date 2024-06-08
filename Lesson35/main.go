package main

import (
	"leetcode-api/config"
	"leetcode-api/handlers"
	"leetcode-api/models"
	"leetcode-api/repository"
	"leetcode-api/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()
	db := config.DB

	db.AutoMigrate(&models.Problem{}, &models.User{}, &models.SolvedProblem{})

	problemRepo := repository.NewProblemRepository(db)
	userRepo := repository.NewUserRepository(db)
	solvedProblemRepo := repository.NewSolvedProblemRepository(db)

	problemService := services.NewProblemService(problemRepo)
	userService := services.NewUserService(userRepo)
	solvedProblemService := services.NewSolvedProblemService(solvedProblemRepo)

	problemHandler := handlers.NewProblemHandler(problemService)
	userHandler := handlers.NewUserHandler(userService)
	solvedProblemHandler := handlers.NewSolvedProblemHandler(solvedProblemService)

	router := mux.NewRouter()

	router.HandleFunc("/problems", problemHandler.GetProblems).Methods("GET")
	router.HandleFunc("/problems/{id}", problemHandler.GetProblem).Methods("GET")
	router.HandleFunc("/problems", problemHandler.CreateProblem).Methods("POST")
	router.HandleFunc("/problems/{id}", problemHandler.UpdateProblem).Methods("PUT")
	router.HandleFunc("/problems/{id}", problemHandler.DeleteProblem).Methods("DELETE")

	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	router.HandleFunc("/solved_problems", solvedProblemHandler.GetSolvedProblems).Methods("GET")
	router.HandleFunc("/solved_problems/{id}", solvedProblemHandler.GetSolvedProblem).Methods("GET")
	router.HandleFunc("/solved_problems", solvedProblemHandler.CreateSolvedProblem).Methods("POST")
	router.HandleFunc("/solved_problems/{id}", solvedProblemHandler.UpdateSolvedProblem).Methods("PUT")
	router.HandleFunc("/solved_problems/{id}", solvedProblemHandler.DeleteSolvedProblem).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
