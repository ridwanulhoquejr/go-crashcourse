package main

import "fmt"

func interfaceType(a Animal) {
	a.Talk()
	// based on the type of the instance, it will call the respective method
	// wheather it is Human or Cat
}

// empty interface
// if we want to pass any type to a function then it would be useful
func emptyInterface(a interface{}) {
	fmt.Println(a)
}

func main() {

	// create a new instance of Human
	h := &Human{}

	// create a new instance of Cat
	c := &Cat{}

	// call the interfaceType function with the Human instance
	// and also with the Cat instance

	// since, both Human and Cat implements the Printer interface
	// we can accept the type which implements the Animal interface

	// so. first we pass the Human instance so it should print "Hello, I am a Human struct type and i can talk"
	// and then we pass the Cat instance so it should print "Hello, I am a Cat struct type but i can't talk"
	interfaceType(h)
	interfaceType(c)

	// calling the emptyInterface function with a string, int and float64
	emptyInterface("Hello, I am a string")
	emptyInterface(10)
	emptyInterface(10.5)

}
