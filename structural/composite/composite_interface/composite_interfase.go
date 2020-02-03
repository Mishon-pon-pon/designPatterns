package composinter

// Athlete ...
type Athlete struct{}

// Train ...
func (a *Athlete) Train() string {
	return "Training!!!"
}

// Swimmer ...
type Swimmer struct {
	Trainig
	Swiming
}

// Animal ...
type Animal struct{}

// Eat ...
func (a *Animal) Eat() string {
	return "Eating!!!"
}

// Shark ...
type Shark struct {
	Eating
	Swiming
}

// Eating ...
type Eating interface {
	Eat() string
}

// Swiming ...
type Swiming interface {
	Swim() string
}

// Trainig ...
type Trainig interface {
	Train() string
}

// SwimmingImpl ...
type SwimmingImpl struct{}

// Swim ...
func (s *SwimmingImpl) Swim() string {
	return "Swiming!!!"
}
