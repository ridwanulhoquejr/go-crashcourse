package main

import "fmt"

type bot interface {

	// this essentially said, whoever you in this program of function of `getGreeting()` and return string,
	// you are now honorary member of a 3rd type `bot` interface
	getGreeting() string
}

type calculation interface {
	getResult(x, y float32) float32
}

type add struct{}
type sub struct{}

type englishBot struct{}
type spanishBot struct{}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi, there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func getAddOrSub(c calculation, x, y float32) {
	fmt.Println("The result of x and y is:", c.getResult(x, y))
	// c has the ability to access both the `getResult` of reciever add and sub
	// because, both the add and sub are a member of calculation interface
}

func (a add) getResult(x, y float32) float32 {
	return x + y
}

func (s sub) getResult(x, y float32) float32 {
	return x - y
}
