package abstract_factory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	GetVehicle(int) (Vehicle, error)
}

// CarFactoryType ...
const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

// GetVehicleFactory ...
func GetVehicleFactory(v int) (VehicleFactory, error) {
	switch v {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Don't have factory with %d number", v))
	}
}
