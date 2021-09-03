package liskovsubstitution

import "testing"

func TestVehicle2(t *testing.T) {

	vehicle2 := &Vehicle2{}
	vehicle2.drive()
	Go2(vehicle2)
}

func TestPoweredVehicle(t *testing.T) {
	poweredVehicle := &PoweredVehicle{}
	poweredVehicle.drive()
	poweredVehicle.startEngine2()

	Go2(poweredVehicle)
}