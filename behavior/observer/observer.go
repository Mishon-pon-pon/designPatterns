package observer

type Observer interface {
	Notify(string)
}

type Publisher struct {
	ObserverList []Observer
}

func (s *Publisher) AddObserver(o Observer) {}

func (s *Publisher) RemoveObserver(o Observer) {}

func (s *Publisher) NotifyObserver(m string) {}


