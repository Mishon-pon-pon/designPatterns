package composite

type Athlete struct{}

func (a *Athlete) Train() string {
	return "Training!"
}

type ComopositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func() string
}

func Swim() string {
	return "Swimming!"
}

type Animal struct{}

func (a *Animal) Eat() string {
	return "Eating!"
}

type Shark struct {
	Animal
	Swim func() string
}
