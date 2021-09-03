package liskovsubstitution

import "fmt"

func Go(vehicle actions) {
	if sled, ok := vehicle.(*Sled); ok {
		sled.pushStart()
	} else {
		vehicle.startEngine()
	}

	vehicle.drive()
}

type actions interface {
	drive()
	startEngine()
}


//type that implements actions interface
type Vehicle struct {
}

//Implements actions
func (v Vehicle) drive() {
	fmt.Println(fmt.Sprintf("Vehicle: %T",v))
	fmt.Println("Vehicle drive")
}

//Implements actions
func (v Vehicle) startEngine() {
	fmt.Println(fmt.Sprintf("Vehicle: %T",v))
	fmt.Println("Vehicle startEngine")
}

//methods
func (v Vehicle) stopEngine() {
	fmt.Println(fmt.Sprintf("Vehicle: %T",v))
	fmt.Println("Vehicle stopEngine")
}

//subtype that composes Vehicle type into struct
//Car doesn't override methods of Vehicle
type Car struct {
	Vehicle
}

//subtype that composes Vehicle type into struct
//Sled overrides methods of Vehicle
type Sled struct {
	Vehicle
}
//Override and implements Actions interface
func (s Sled) startEngine() {
	fmt.Println(fmt.Sprintf("Sled: %T",s))
	fmt.Println("Sled startEngine")
}

//Override of Vehicle method
func (s Sled) stopEngine() {
	fmt.Println(fmt.Sprintf("Sled: %T",s))
	fmt.Println("Sled stopEngine")
}

//Sled method only
func (s Sled) pushStart() {
	fmt.Println(fmt.Sprintf("Sled: %T",s))
	fmt.Println("Sled pushStart")
}

/*
Why violating LSP?
- Sled has "pushStart" method which can't be triggered by Car

 */