package abstract_factory

type CruiseMotorbike struct{}

func (m *CruiseMotorbike) GetWheels() int {
	return 2
}

func (m *CruiseMotorbike) GetSeats() int {
	return 2
}
