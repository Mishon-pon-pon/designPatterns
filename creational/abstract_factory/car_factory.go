package abstract_factory

import (
	"errors"
	"fmt"
)

// LuxuryCar ...
const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

// CarFactory ...
type CarFactory struct{}

// NewVehicle ...
func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {

	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))

	}
}
