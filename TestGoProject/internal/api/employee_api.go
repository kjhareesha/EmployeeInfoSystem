package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"TestGoProject/internal/model"
	"TestGoProject/internal/service"

	"github.com/gorilla/mux"
)

type EmployeeHandler struct {
	Service *service.EmployeeService
}

func RegisterEmployeeRoutes(r *mux.Router, service *service.EmployeeService) {
	handler := &EmployeeHandler{Service: service}
	r.HandleFunc("/employees", handler.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees", handler.GetAllEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", handler.GetEmployeeByID).Methods("GET")
	r.HandleFunc("/employees/{id}", handler.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", handler.DeleteEmployee).Methods("DELETE")
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	emp := &model.Employee{}
	if err := json.NewDecoder(r.Body).Decode(emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateEmployee(emp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func (h *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.Service.GetAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(employees)
}

func (h *EmployeeHandler) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	emp, err := h.Service.GetEmployeeByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(emp)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	emp := &model.Employee{ID: id}
	if err := json.NewDecoder(r.Body).Decode(emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.UpdateEmployee(emp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(emp)
}

func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.DeleteEmployee(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
