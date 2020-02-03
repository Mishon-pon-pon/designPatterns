package abstract_factory

// CruiseMotobike ...
type CruiseMotobike struct{}

// NumWheels ...
func (c *CruiseMotobike) NumWheels() int {
	return 2
}

// NumSeats ...
func (c *CruiseMotobike) NumSeats() int {
	return 1
}

// GetMotobikeType ...
func (c *CruiseMotobike) GetMotobikeType() int {
	return CruiseMotobikeType
}
