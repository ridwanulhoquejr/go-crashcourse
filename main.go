package main

import (
	"fmt"
)

// // packages == project == workspace
// // we only have one package per project

// // there are two types of packages:
// // 1. Executable
// // 2. Reusable

// // Executable package is used to create an executable file (.exe)
// // Reusable package is used to put reusable code in one place and use it in other packages like dependencies

// func main() {

// 	// slice
// 	// cards := newDeck()
// 	// cards.shuffle()
// 	// cards.saveToFile("my_cards")

// 	// cards := newDeckFromFile("my_cards")

// 	// cards.print()

// 	// creating a struct object, this is
// 	// ridwan := person{firstName: "Ridwan", lastName: "Hoque",
// 	// 	contact: contactInfo{
// 	// 		email:   "ridwan@gmail.com",
// 	// 		zipcode: 4203,
// 	// 	},
// 	// }

// 	// ridwan.updatePerson("Rid1")
// 	// // var huma person
// 	// ridwan.printPerson()
// 	// // huma.printPerson()

// 	// // fmt.Println(ridwan)
// 	// // fmt.Println(huma)
// 	// // fmt.Printf("%+v", ridwan)

// 	// // initializing a map
// 	// // var clr = color{
// 	// // 	"red":   "#ff0000",
// 	// // 	"green": "#4bf745",
// 	// // 	"white": "#ffffff",
// 	// // }

// 	// // or an empty map
// 	// clr := make(color)
// 	// clr["red"] = "#ff0000"
// 	// clr["green"] = "#4bf745"
// 	// clr.print()

// 	// delete(clr, "red")
// 	// clr.print()

// 	// /// interfaces

// 	// eb := englishBot{}
// 	// sb := spanishBot{}
// 	// printGreeting(eb)
// 	// printGreeting(sb)

// 	// a := add{}
// 	// s := sub{}

// 	// getAddOrSub(a, 10.5, 5.5)
// 	// getAddOrSub(s, 10.5, 5.5)

// 	/// HTTP request
// 	res, err := http.Get("http://www.google.com")

// 	if err != nil {
// 		println("Error:", err)
// 		os.Exit(1)
// 	}

// 	lw := logWriter{}

// 	io.Copy(lw, res.Body)

// 	// bs := make([]byte, 99999)
// 	// res.Body.Read(bs)
// 	// fmt.Println(string(bs))

// }

// array: fixed length list of things
// slice: an array that can grow or shrink

func main() {

	// intSlice := []int{1, 2, 3}
	// fmt.Println(sumSlice[int](intSlice))

	/// generic as a struct
	// var gasCasr = car[gasEngine]{
	// 	carMake:  "Honda",
	// 	carModel: "2024",
	// 	engine: gasEngine{
	// 		gallons: 11,
	// 		mpg:     33,
	// 	},
	// }
	// fmt.Println(gasCasr)

	// var electricCar = car[electricEngine]{
	// 	carMake:  "Honda",
	// 	carModel: "2024",
	// 	engine: electricEngine{
	// 		kwh:   9,
	// 		mpkwh: 77,
	// 	},
	// }

	// fmt.Println(electricCar)

	var p *int = new(int) // p will holds a free memory location means, points to a mem address
	// keep in my that, pointer variable only hold another memory address while he also does has one address.

	fmt.Println("P vlaue and initial value ", p, &p, *p)

	*p = 10 // change the value of p which points to a mem address with value of 10

	fmt.Println("Initial vlaue, p points to add and current value: ", p, &p, *p)
	*p = 12          // change the value of p which points to a mem address with value of 10
	*(&p) = new(int) // this will change the memory address of what p points to
	fmt.Println("Initial vlaue, p points to add and current value: ", p, &p, *p)

	// var i int = 9

	// fmt.Println("Value of the i and memory address:", i, &i)

	// test commit
}
