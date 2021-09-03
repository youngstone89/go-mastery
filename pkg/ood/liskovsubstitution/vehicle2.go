package liskovsubstitution

import "fmt"

func Go2(vehicle actions2) {
	switch concrete := vehicle.(type) {
	case poweredActions:
		concrete.startEngine2()
	case unpoweredActions:
		concrete.pushStart2()
	}

	vehicle.drive()
}

type actions2 interface {
	drive()
}

//interface that composes actions2 interface
type poweredActions interface {
	actions2
	startEngine2()
	stopEngine2()
}

type unpoweredActions interface {
	actions2
	pushStart2()
}

type Vehicle2 struct {
}

//implements actions2 interface
func (v Vehicle2) drive() {
	fmt.Println(fmt.Sprintf("drive : %T",v))
}

//subtype of Vehicle2
type PoweredVehicle struct {
	Vehicle2
}

//implements poweredActions interface
func (v PoweredVehicle) startEngine2() {
	fmt.Println(fmt.Sprintf("startEngine2 : %T",v))
}

//subtype of PoweredVehicle
type Car2 struct {
	PoweredVehicle
}

//subtype of Vehicle2
type Buggy struct {
	Vehicle2
}

func (b Buggy) pushStart() {
	fmt.Println(fmt.Sprintf("startEngine2 : %T",b))
}