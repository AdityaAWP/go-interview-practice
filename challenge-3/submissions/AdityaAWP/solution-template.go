package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
	m.Employees = append(m.Employees, e)
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
	for i, k := range m.Employees {
		if k.ID == id {
			m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
		}
	}
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
	totalEmployee := len(m.Employees)
	totalSalary := 0.0
	if totalEmployee == 0 {
		return 0.0
	}
	for _, v := range m.Employees {
		totalSalary += v.Salary
	}
	averageSalary := totalSalary / float64(totalEmployee)
	return averageSalary
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
	for i, v := range m.Employees {
		if m.Employees[i].ID == id {
			fmt.Println(v.ID)
			return &m.Employees[i]
		}
	}

	return nil
}

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
	manager.RemoveEmployee(1)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average Salary: %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}
}
