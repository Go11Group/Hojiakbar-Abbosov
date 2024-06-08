package handlers

import (
	"encoding/json"
	"leetcode-api/models"
	"leetcode-api/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProblemHandler struct {
	Service *services.ProblemService
}

func NewProblemHandler(service *services.ProblemService) *ProblemHandler {
	return &ProblemHandler{Service: service}
}

func (h *ProblemHandler) GetProblems(w http.ResponseWriter, r *http.Request) {
	problems, err := h.Service.GetAllProblems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(problems)
}

func (h *ProblemHandler) GetProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	problem, err := h.Service.GetProblemByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(problem)
}

func (h *ProblemHandler) CreateProblem(w http.ResponseWriter, r *http.Request) {
	var problem models.Problem
	_ = json.NewDecoder(r.Body).Decode(&problem)
	problem, err := h.Service.CreateProblem(problem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(problem)
}

func (h *ProblemHandler) UpdateProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var problem models.Problem
	_ = json.NewDecoder(r.Body).Decode(&problem)
	problem.ID = uint(id)
	problem, err := h.Service.UpdateProblem(problem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(problem)
}

func (h *ProblemHandler) DeleteProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	err := h.Service.DeleteProblem(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
