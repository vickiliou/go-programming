package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("hello!")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "vic",
	}

	saySomething(&p1)

	p1.speak()
}
