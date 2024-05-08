package main

import "fmt"

func main() {

	// pointer

	var x int = 10
	var ptr *int = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)

	fmt.Println("addresse of ptr points to:", ptr)
	fmt.Println("Value of ptr that points to:", *ptr)
	fmt.Println("Address of ptr:", &ptr)

	//
	y := &Str{number: 10}
	z := 10
	fmt.Println("Value of y:", y.number)
	fmt.Println("Value of z:", z)
	y.addOneWithStruct()
	addOneWithParams(&z)
	fmt.Println("Value of y after addOne:", y.number)
	fmt.Println("Value of z after addOne:", z)

}
