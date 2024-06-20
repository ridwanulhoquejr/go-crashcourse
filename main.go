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

/// Anothr good example:
// package main

// import (
// 	"fmt"
// 	"log"
// 	"strconv"
// )

// // ? to implement this interface, any GO types must have a method String() which returns a string
// type Stringer interface {
// 	String() string
// }

// type Book struct {
// 	Title  string
// 	Author string
// 	Sold   Stringer
// }

// type Count int

// // now we can say that Book implements the Stringer interface
// func (b Book) String() string {
// 	return fmt.Sprintf("Book: %s by %s and Sold %s", b.Title, b.Author, b.Sold)
// }

// // now we can say that Count implements the Stringer interface
// func (c Count) String() string {
// 	return strconv.Itoa(int(c))
// }

// // ? so both Book and Count implement the Stringer interface,
// // ? now we can pass both Book and Count objects to WriteLog as arguments
// func WriteLog(s Stringer) {
// 	// at this moment, we don't know if s is a Book or a Count
// 	// but we know that it implements the Stringer interface
// 	// So in runtime, it will be either a Book or a Count and it will call the respective String() method
// 	log.Print(s.String())
// }

// func main() {

// 	// ? now we can pass Book and Count to WriteLog
// 	b := Book{
// 		Title:  "The Art of Computer Programming",
// 		Author: "Donald Knuth",
// 		// Sold:   Book{},
// 		Sold: Count(10),
// 	}
// 	WriteLog(b)

// 	c := Count(3)
// 	WriteLog(c)

// 	//! if we try to pass a type that does not implement the Stringer interface, we will get a compile-time error
// 	// i := 3
// 	// WriteLog(i)
// }

