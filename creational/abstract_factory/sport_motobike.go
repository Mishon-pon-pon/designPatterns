package abstract_factory

// SportMotobike ...
type SportMotobike struct{}

// NumWheels ...
func (s *SportMotobike) NumWheels() int {
	return 2
}

// NumSeats ...
func (s *SportMotobike) NumSeats() int {
	return 1
}

// GtMotobikeType ...
func (s *SportMotobike) GetMotobikeType() int {
	return SportMotobikeType
}
