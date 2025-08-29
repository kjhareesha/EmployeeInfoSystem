package service

import (
	"d:/POC/TestGoProject/internal/model"
	"d:/POC/TestGoProject/internal/repository"
	"fmt"
)

type EmployeeService struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{Repo: repo}
}

func (s *EmployeeService) CreateEmployee(emp *model.Employee) error {
	// Validate required fields
	if emp.FirstName == "" || emp.LastName == "" || emp.Email == "" || emp.Position == "" {
		return fmt.Errorf("all fields are required")
	}

	// Validate allowed positions
	allowedPositions := map[string]bool{"Manager": true, "Developer": true, "HR": true, "Sales": true}
	if !allowedPositions[emp.Position] {
		return fmt.Errorf("invalid position: %s", emp.Position)
	}

	// Check for duplicate email
	employees, err := s.Repo.GetAll()
	if err != nil {
		return fmt.Errorf("error checking existing employees: %w", err)
	}
	for _, e := range employees {
		if e.Email == emp.Email {
			return fmt.Errorf("email already exists")
		}
	}

	return s.Repo.Create(emp)
}

func (s *EmployeeService) GetEmployeeByID(id int) (*model.Employee, error) {
	return s.Repo.GetByID(id)
}

func (s *EmployeeService) GetAllEmployees() ([]model.Employee, error) {
	return s.Repo.GetAll()
}

func (s *EmployeeService) UpdateEmployee(emp *model.Employee) error {
	// Validate required fields
	if emp.FirstName == "" || emp.LastName == "" || emp.Email == "" || emp.Position == "" {
		return fmt.Errorf("all fields are required")
	}

	// Validate allowed positions
	allowedPositions := map[string]bool{"Manager": true, "Developer": true, "HR": true, "Sales": true}
	if !allowedPositions[emp.Position] {
		return fmt.Errorf("invalid position: %s", emp.Position)
	}

	// Check for duplicate email (excluding current employee)
	employees, err := s.Repo.GetAll()
	if err != nil {
		return fmt.Errorf("error checking existing employees: %w", err)
	}
	for _, e := range employees {
		if e.Email == emp.Email && e.ID != emp.ID {
			return fmt.Errorf("email already exists")
		}
	}

	return s.Repo.Update(emp)
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	return s.Repo.Delete(id)
}
