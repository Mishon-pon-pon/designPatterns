package abstract_factory

import (
	"errors"
	"fmt"
)

// LuxuryCar ...
const (
	SportMotobikeType  = 1
	CruiseMotobikeType = 2
)

// MotobikeFactory ...
type MotobikeFactory struct{}

// Build ...
func (m *MotobikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotobikeType:
		return new(SportMotobike), nil
	case CruiseMotobikeType:
		return new(CruiseMotobike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}
