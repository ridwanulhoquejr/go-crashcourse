package main

// declare an interface called Printer
// so any type (in our case it is, struct) which tries
// to implement this interface should have a method called Print
type Animal interface {
	Talk()
}

// so we declare a struct called Person
// and try to implement the Printer interface
type Human struct {
}

type Cat struct{}

// ! NOTE: in go there is no explicit declaration of implementing an interface
// so we just have to declare a method with the same name as the interface
// and that method should have a receiver of the same type
// so here we have a method called Print with a receiver of type Person

// now we can say that Person [struct] implements the Printer interface
func (p *Human) Talk() {
	println("Hello, I am a Human struct type and i can talk")
}

func (p *Cat) Talk() {
	println("Hello, I am a Cat struct type but i can't talk")
}
