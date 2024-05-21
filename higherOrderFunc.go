package main

import "fmt"

// Function as Argument:

func applyOperations(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func add(x, y int) int {
	return x + y
}

// return a function
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// both takes and returns a function
// return a function:
func makeGreeter(greeting string) func(string) string {
	return func(name string) string {
		return greeting + " " + name
	}
}

// now take the above function as an parameter
// Function as Argument:
func greet(name string, greeter func(string) string) {
	fmt.Println(greeter(name))
}

// take a list of int and process only even numbers
func evenNumbers(n []int) func() []int {
	var even []int

	return func() []int {
		for _, v := range n {
			if v%2 == 0 {
				even = append(even, v)
			}
		}
		return even
	}
}

// function as an argument to filter even numbers
func evenNumbers2(n []int, f func([]int) []int) []int {
	return f(n)
}
