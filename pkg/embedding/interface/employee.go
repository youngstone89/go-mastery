package infembedding

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func NewEmployee(name string, ID string) *Employee {
	return &Employee{Name: name, ID: ID}
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func NewManager(employee Employee, reports []Employee) *Manager {
	return &Manager{Employee: employee, Reports: reports}
}

func (m Manager) FindNewEmployees() []Employee {
	// do business logic
	return m.Reports
}
