package abstract_factory

type FamilyCar struct{}

func (c *FamilyCar) GetWheels() int {
	return 4
}

func (c *FamilyCar) GetSeats() int {
	return 5
}
