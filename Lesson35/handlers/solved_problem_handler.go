package handlers

import (
	"encoding/json"
	"leetcode-api/models"
	"leetcode-api/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SolvedProblemHandler struct {
	Service *services.SolvedProblemService
}

func NewSolvedProblemHandler(service *services.SolvedProblemService) *SolvedProblemHandler {
	return &SolvedProblemHandler{Service: service}
}

func (h *SolvedProblemHandler) GetSolvedProblems(w http.ResponseWriter, r *http.Request) {
	solvedProblems, err := h.Service.GetAllSolvedProblems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(solvedProblems)
}

func (h *SolvedProblemHandler) GetSolvedProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	solvedProblem, err := h.Service.GetSolvedProblemByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(solvedProblem)
}

func (h *SolvedProblemHandler) CreateSolvedProblem(w http.ResponseWriter, r *http.Request) {
	var solvedProblem models.SolvedProblem
	_ = json.NewDecoder(r.Body).Decode(&solvedProblem)
	solvedProblem, err := h.Service.CreateSolvedProblem(solvedProblem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(solvedProblem)
}

func (h *SolvedProblemHandler) UpdateSolvedProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var solvedProblem models.SolvedProblem
	_ = json.NewDecoder(r.Body).Decode(&solvedProblem)
	solvedProblem.ID = uint(id)
	solvedProblem, err := h.Service.UpdateSolvedProblem(solvedProblem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(solvedProblem)
}

func (h *SolvedProblemHandler) DeleteSolvedProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	err := h.Service.DeleteSolvedProblem(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
