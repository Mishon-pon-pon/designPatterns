package abstract_factory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := GetVehicleFactory(MotobikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.NewVehicle(SportMotobikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels and %d seats\n", motorbikeVehicle.NumWheels(), motorbikeVehicle.NumSeats())

	sportBike, ok := motorbikeVehicle.(Motobike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotobikeType())

	motorbikeVehicle, err = motorbikeF.NewVehicle(CruiseMotobikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.NumWheels())

	cruiseBike, ok := motorbikeVehicle.(Motobike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Cruise motorbike has type %d\n", cruiseBike.GetMotobikeType())

	motorbikeVehicle, err = motorbikeF.NewVehicle(3)
	if err == nil {
		t.Fatal("Motorbike of type 3 should not be recognized")
	}
}

func TestCarFactory(t *testing.T) {
	carF, err := GetVehicleFactory(3)
	if err == nil {
		t.Fatal("Car factory with id 3 should not be recognized")
	}

	carF, err = GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d seats and %d wheels\n", carVehicle.NumSeats(), carVehicle.NumWheels())

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())

	carVehicle, err = carF.NewVehicle(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d seats\n", carVehicle.NumWheels())

	familiarCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Familiar car has %d doors.\n", familiarCar.NumDoors())

	carVehicle, err = carF.NewVehicle(3)
	if err == nil {
		t.Fatal("Car of type 3 should not be recognized")
	}
}
