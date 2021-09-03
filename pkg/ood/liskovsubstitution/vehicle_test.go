package liskovsubstitution

import "testing"

func TestVehicle(t *testing.T) {
	vehicle := &Vehicle{}
	vehicle.startEngine()
	vehicle.stopEngine()
	vehicle.drive()
	Go(vehicle)
}

func TestSled(t *testing.T) {
	sled := &Sled{}
	sled.startEngine()
	sled.drive()
	sled.stopEngine()
	Go(sled)
}

func TestCar(t *testing.T) {
	car := &Car{}
	car.startEngine()
	car.stopEngine()
	car.drive()
	Go(car)
}
