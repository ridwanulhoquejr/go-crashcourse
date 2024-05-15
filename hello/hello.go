package main

import (
	"fmt"

	"github.com/ridwanulhoquejr/go-crashcourse/greetings"
	"github.com/ridwanulhoquejr/go-crashcourse/hello/subhello"
	"github.com/ridwanulhoquejr/go-crashcourse/talking"
)

func main() {

	message := greetings.Hello("Gladys")
	fmt.Println(message)

	// get the message from the talking package
	talk := talking.Talk("Ridwan")
	fmt.Println(talk)

	fmt.Println(subhello.FromSubHello("Ridwan"))

}
