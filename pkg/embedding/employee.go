package embedding

import "fmt"

func employeeTest()  {

	e := NewEmployee("YEONGSEOKKIM","1")
	e1 := NewEmployee("YEONGSEOKKIM","2")
	e2 := NewEmployee("YEONGSEOKKIM","3")
	e3 := NewEmployee("YEONGSEOKKIM","4")
	e4 := NewEmployee("YEONGSEOKKIM","5")
	reports := make([]Employee,0)
	//reports = append(reports, *e)
	reports = append(reports, *e1)
	reports = append(reports, *e2)
	reports = append(reports, *e3)
	reports = append(reports, *e4)
	m := NewManager(*e, reports)

	for _,v := range m.FindNewEmployees() {
		fmt.Println(v.Description())
	}
}

type Employee struct {
	Name         string
	ID           string
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