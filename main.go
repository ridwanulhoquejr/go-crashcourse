package main

import (
	"io"
	"net/http"
	"os"
)

// packages == project == workspace
// we only have one package per project

// there are two types of packages:
// 1. Executable
// 2. Reusable

// Executable package is used to create an executable file (.exe)
// Reusable package is used to put reusable code in one place and use it in other packages like dependencies

func main() {

	// slice
	// cards := newDeck()
	// cards.shuffle()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")

	// cards.print()

	// creating a struct object, this is
	// ridwan := person{firstName: "Ridwan", lastName: "Hoque",
	// 	contact: contactInfo{
	// 		email:   "ridwan@gmail.com",
	// 		zipcode: 4203,
	// 	},
	// }

	// ridwan.updatePerson("Rid1")
	// // var huma person
	// ridwan.printPerson()
	// // huma.printPerson()

	// // fmt.Println(ridwan)
	// // fmt.Println(huma)
	// // fmt.Printf("%+v", ridwan)

	// // initializing a map
	// // var clr = color{
	// // 	"red":   "#ff0000",
	// // 	"green": "#4bf745",
	// // 	"white": "#ffffff",
	// // }

	// // or an empty map
	// clr := make(color)
	// clr["red"] = "#ff0000"
	// clr["green"] = "#4bf745"
	// clr.print()

	// delete(clr, "red")
	// clr.print()

	// /// interfaces

	// eb := englishBot{}
	// sb := spanishBot{}
	// printGreeting(eb)
	// printGreeting(sb)

	// a := add{}
	// s := sub{}

	// getAddOrSub(a, 10.5, 5.5)
	// getAddOrSub(s, 10.5, 5.5)

	/// HTTP request
	res, err := http.Get("http://www.google.com")

	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, res.Body)

	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

}

// array: fixed length list of things
// slice: an array that can grow or shrink
