package composinter

import "testing"

func TestComposinter(t *testing.T) {
	swimmer := Swimmer{
		&Athlete{},
		&SwimmingImpl{},
	}

	var str string = swimmer.Train()
	t.Log(str)
	str = swimmer.Swim()
	t.Log(str)

	var shark Shark = Shark{
		&Animal{},
		&SwimmingImpl{},
	}

	str = shark.Eat()
	t.Log(str)
	str = shark.Swim()
	t.Log(str)
}
