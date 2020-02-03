package composite

import "testing"

func TestComposite(t *testing.T) {
	localSwim := Swim
	swimmer := ComopositeSwimmerA{
		MySwim: localSwim,
	}

	var str string = swimmer.MyAthlete.Train()
	t.Log(str)
	str = swimmer.MySwim()
	t.Log(str)

	fish := Shark{
		Swim: localSwim,
	}

	str = fish.Eat()
	t.Log(str)
	str = fish.Swim()
	t.Log(str)
}
