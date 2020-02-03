package composite

// Athlete ...
type Athlete struct{}

// Train ...
func (a *Athlete) Train() string {
	return "Training!"
}

// ComopositeSwimmerA ...
type ComopositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func() string
}

// Swim ...
func Swim() string {
	return "Swimming!"
}

// Animal ...
type Animal struct{}

// Eat ...
func (a *Animal) Eat() string {
	return "Eating!"
}

// Shark ...
type Shark struct {
	Animal
	Swim func() string
}
