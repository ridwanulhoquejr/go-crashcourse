package main

func interfaceType(a Animal) {
	a.Talk()
}

func main() {

	// create a new instance of Human
	h := Human{}

	// create a new instance of Cat
	c := Cat{}

	// call the interfaceType function with the Human instance
	// and also with the Cat instance

	// since, both Human and Cat implements the Printer interface
	// we can accept the type which implements the Animal interface

	// so. first we pass the Human instance so it should print "Hello, I am a Human struct type"
	// and then we pass the Cat instance so it should print "Hello, I am a Cat struct type"
	interfaceType(&h)
	interfaceType(&c)

}
