package main

import "fmt"

func handlePanic() {
	a := recover()

	if a != nil {
		fmt.Println("Recovered from panic")
	}
}

func main() {

	// defer, panic and recover
	//
	// Defer is used to ensure that a function call is performed later in a program’s execution,
	// usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.

	// panic, recover and defer

	fmt.Println("Before panic")
	runPanic()
	fmt.Println("After panic recovery")

}

// however we can recover from this panic, by using recover function and defer inside the panic function
func runPanic() {
	defer handlePanic()
	panic("Panic stated, so below code will not be executed")
}

// Bottom line:
// 1. A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns.
// 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
// 3. It is similar to the finally block in Other languages.
