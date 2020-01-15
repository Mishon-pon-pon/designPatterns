package creational

// BuildProcess ...
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// ManufacturingDirector ...
type ManufacturingDirector struct {
	builder BuildProcess
}

// Construct ...
func (m *ManufacturingDirector) Construct() {
	m.builder.SetSeats().SetStructure().SetWheels()
}

// SetBuilder ...
func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

// VehicleProduct ...
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// CarBuilder ...
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels ...
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetSeats ...
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetStructure ...
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

// GetVehicle ...
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// BikeBuilder ...
type BikeBuilder struct {
	v VehicleProduct
}

// SetWheels ...
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

// SetSeats ...
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

// SetStructure ...
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motobike"
	return b
}

// GetVehicle ...
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// BusBuilder ...
type BusBuilder struct {
	v VehicleProduct
}

// SetWheels ...
func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 6
	return b
}

// SetSeats ...
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 24
	return b
}

// SetStructure ...
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

// GetVehicle ...
func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}
