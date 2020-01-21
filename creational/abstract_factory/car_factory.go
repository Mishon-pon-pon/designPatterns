package abstract_factory

import (
	"errors"
	"fmt"
)

type CarFactory struct{}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Don't have type with %d number", v))
	}
}
