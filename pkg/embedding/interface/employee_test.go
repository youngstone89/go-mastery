package infembedding_test

import (
	"fmt"
	inf "go-mastery/pkg/embedding/interface"
	"testing"
)

func TestManager(t *testing.T) {

	e := inf.NewEmployee("MANAGER-YEONGSEOKKIM", "1")
	e1 := inf.NewEmployee("YEONGSEOKKIM", "2")
	e2 := inf.NewEmployee("YEONGSEOKKIM", "3")
	e3 := inf.NewEmployee("YEONGSEOKKIM", "4")
	e4 := inf.NewEmployee("YEONGSEOKKIM", "5")
	reports := make([]inf.Employee, 0)
	//reports = append(reports, *e)
	reports = append(reports, *e1)
	reports = append(reports, *e2)
	reports = append(reports, *e3)
	reports = append(reports, *e4)
	m := inf.NewManager(*e, reports)

	fmt.Println(m.Description())
	for _, v := range m.FindNewEmployees() {
		fmt.Println(v.Description())
	}

}
