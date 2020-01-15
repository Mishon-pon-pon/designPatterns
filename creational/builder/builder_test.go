package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Where on a car must be 4 wheels and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be Car and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they where %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motobike := bikeBuilder.GetVehicle()
	// motobike.Seats = 1

	if motobike.Wheels != 2 {
		t.Errorf("Wheels on a motobike must be 2 and they where %d\n", motobike.Wheels)
	}

	if motobike.Seats != 2 {
		t.Errorf("Seats on a motobike must be 2 and they were %d\n", motobike.Seats)
	}

	if motobike.Structure != "Motobike" {
		t.Errorf("Structure on motorbike must be Motobike  and was %s\n", motobike.Structure)
	}

	busBuilder := &BusBuilder{}
	manufacturingComplex.SetBuilder(busBuilder)
	manufacturingComplex.Construct()

	bus := busBuilder.GetVehicle()

	if bus.Wheels != 6 {
		t.Errorf("Wheels on a bus must be 6 and they were %d\n", bus.Wheels)
	}

	if bus.Seats != 24 {
		t.Errorf("Seats on a bus must be 24 and they were %d\n", bus.Seats)
	}

	if bus.Structure != "Bus" {
		t.Errorf("Structure on a bus must be Bus and were %s\n", bus.Structure)
	}

}
