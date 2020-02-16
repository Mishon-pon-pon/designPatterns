package main

import "fmt"

var observer *Observer

func GetObserver() *Observer {
	if observer == nil {
		observer = &Observer{}
	}
	return observer
}

type Observer struct {
	Store map[string] []CallBack
}

func (o *Observer) Subscribe(event string, f CallBack) {
	if o.Store == nil{
		o.Store = make(map[string][]CallBack)
	}
	if _, ok := o.Store[event]; ok {
		o.Store[event] = append(o.Store[event], f)
	} else {
		o.Store[event] = []CallBack{f}
	}
}

func (o *Observer) Publish(event string, data interface{}) {
	if _, ok := o.Store[event]; ok {
		for _, f := range o.Store[event] {
			f(data)
		}
	} else {
		return
	}
}

type CallBack func(interface{})

func main() {
	f1 := func(data interface{}){
		fmt.Println("Hello I am function f1 and my data is -", data)
	}
	f2 := func(data interface{}) {
		fmt.Println("Hello I'm function f2 and my data is -", data)
	}

	var o *Observer = GetObserver()
	var o2 *Observer = GetObserver()

	o.Subscribe("Event1", f2)
	o2.Subscribe("Event3", func(data interface{}) {
		fmt.Println("Hello I'm anonymous function and my data is -", data)
	})
	o2.Subscribe("Event1", f1)
	o2.Subscribe("Event2", f2)

	fmt.Println("All function of 'Event1'")
	o.Publish("Event1", "Data of Event1")

	fmt.Println("All function of 'Event2'")
	o.Publish("Event2", struct {
		message string
	}{message: "Struct of Event2"})

	fmt.Println("All function of 'Event3'")
	o2.Publish("Event3", "Data of Event3")
}
