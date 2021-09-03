package liskovsubstitution

import "fmt"

func Go3(vehicle actions3) {
	vehicle.start()
	vehicle.drive()

}

type actions3 interface {
	start()
	drive()
}


//subtype of poweredVehicle3 & type that implements actions3 interface
type Car3 struct {
	poweredVehicle3
}

//executed composed type's method
func (c Car3) start() {
	c.poweredVehicle3.startEngine()
}

func (c Car3) drive() {
	fmt.Println(fmt.Sprintf("drive : %T",c))
}

//type that doesn't actions3 interface
type poweredVehicle3 struct {
}
//poweredVehicle3 method only
func (p poweredVehicle3) startEngine() {
	fmt.Println(fmt.Sprintf("start : %T",p))
}

//type that implements actions3 interface
type Buggy3 struct {
}

//implements actions3 start() interface method
func (b Buggy3) start() {
	fmt.Println(fmt.Sprintf("start : %T",b))
}

//implements actions3 drive() interface method
func (b Buggy3) drive() {
	fmt.Println(fmt.Sprintf("drive : %T",b))
}