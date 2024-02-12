package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ? deck is a slice of strings
type deck []string

func newDeck() deck {

	// first create an empty deck, then add cards slices then ran 2 for loop to add cards

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// this is a receiver function, whcih belongs to the type deck
// so, we can access this function who has a deck type
// isnt it like dart extension?? yes, it is

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// this will return two values, one is hand and another is remaining cards
func deal(d deck, handSize int) (deck, deck) {
	// here (deck, deck) is the return type, its the way to return multiple values from a function

	// and [:] is behaves like python slicing
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	// type conversion: from deck to []string and then join the whole string with comma
	return strings.Join([]string(d), ",")
}

// saveToFile function will save the deck to a file
func (d deck) saveToFile(filename string) error {

	// so, to save a file to our local machine, we have to first convert the []deck to a []string, then convert it to []byte
	// then os.WriteFile

	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

// readFromFile function
func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)

	if err != nil {
		// if there is an error, then print the error and exit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// string(byteSlice) will convert the byteSlice to a string
	// then strings.Split will split the string by comma and return a slice of strings
	// then we will convert the slice of strings to a deck and return it
	return deck(strings.Split(string(byteSlice), ","))
}

// shuffle function
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
