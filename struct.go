package main

import "fmt"

// struct is similar to a class. where all the properties of the class consists
// we can use embedding struct without nam cotact
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipcode int
}

func (p *person) updatePerson(newName string) {

	// *person in reciever func tells this func that we deal with a pointer of person
	// essentially by telling that we have to deal with same struct
	(*p).firstName = newName
}

func (p person) printPerson() {
	fmt.Printf("%+v", p)
}
