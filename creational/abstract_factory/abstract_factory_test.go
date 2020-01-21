package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	af, err := GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fail()
	}

	cf, err := af.GetVehicle(LuxuryCarType)
	if err != nil {
		t.Fail()
	}

	if cf.GetSeats() != 4 {
		t.Error("Luxury car must have 4 seats")
	}

	cf, err = af.GetVehicle(FamilyCarType)
	if err != nil {
		t.Fail()
	}

	if cf.GetSeats() != 5 {
		t.Error("Family car must have 5 seats")
	}

	af, err = GetVehicleFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fail()
	}

	mf, err := af.GetVehicle(SportMotorbikeType)
	if err != nil {
		t.Fail()
	}

	if mf.GetSeats() != 1 {
		t.Error("Sport motorbike must have 1 seats")
	}

	mf, err = af.GetVehicle(CruiseMotorbikeType)
	if err != nil {
		t.Fail()
	}

	if mf.GetSeats() != 2 {
		t.Error("Cruise motorbike must have 2 seats")
	}
}
