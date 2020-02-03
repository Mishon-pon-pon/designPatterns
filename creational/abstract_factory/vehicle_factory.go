package abstract_factory

import (
	"errors"
	"fmt"
)

// VehicleFactory ...
type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

// CarFactoryType ...
const (
	CarFactoryType      = 1
	MotobikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotobikeFactoryType:
		return new(MotobikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n", f))
	}
}
