package repository

import (
	"d:/POC/TestGoProject/internal/model"
	"database/sql"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Create(employee *model.Employee) error {
	query := "INSERT INTO employees (first_name, last_name, email, position) VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, employee.FirstName, employee.LastName, employee.Email, employee.Position)
	return err
}

func (r *EmployeeRepository) GetByID(id int) (*model.Employee, error) {
	query := "SELECT id, first_name, last_name, email, position FROM employees WHERE id = ?"
	row := r.DB.QueryRow(query, id)
	emp := &model.Employee{}
	if err := row.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Email, &emp.Position); err != nil {
		return nil, err
	}
	return emp, nil
}

func (r *EmployeeRepository) GetAll() ([]model.Employee, error) {
	query := "SELECT id, first_name, last_name, email, position FROM employees"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []model.Employee{}
	for rows.Next() {
		emp := model.Employee{}
		if err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Email, &emp.Position); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

func (r *EmployeeRepository) Update(employee *model.Employee) error {
	query := "UPDATE employees SET first_name=?, last_name=?, email=?, position=? WHERE id=?"
	_, err := r.DB.Exec(query, employee.FirstName, employee.LastName, employee.Email, employee.Position, employee.ID)
	return err
}

func (r *EmployeeRepository) Delete(id int) error {
	query := "DELETE FROM employees WHERE id=?"
	_, err := r.DB.Exec(query, id)
	return err
}
