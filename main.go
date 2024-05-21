package main

import "fmt"

//? Closures: Its a function that references variables from outside its body.
// in this case, the variable sum is referenced by the closure function returned by adder.

// adder returns a closure that captures the sum variable.
func adder() func(int) int {
	sum := 0 // sum is a variable in the scope of the adder function.

	// This anonymous function (closure) captures the sum variable.
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	// example of closures:
	pos, neg := adder(), adder()

	for i := 0; i <= 10; i++ {
		fmt.Println(
			pos(i),
			neg(-i),
		)
	}

	///
	// Characteristics of Higher Order Functions:
	// 1. Accepts functions as parameters
	// 2. Returns a function -> this is clousure
	// 3. both 1 and 2
	//
	// example of higher order functions: 1 -> function as an argument
	sum := applyOperations(10, 20, add)
	fmt.Println(sum)

	// or we can direct wrote a func as an argument
	sub := applyOperations(20, 10, func(a, b int) int {
		return a - b
	})
	fmt.Println(sub)

	// example of higher order functions: 2 -> return a function
	multiplyBy2 := multiplier(2)
	fmt.Println(multiplyBy2(10))

	// example of higher order functions: 3 -> both takes and returns a function
	helloGreeter := makeGreeter("Hello") // closure
	greet("World", helloGreeter)         // we execute the closure in the greet function

	// exercise example:
	// Filter even numbers using -> return a function
	en := evenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(en())

	// Filter even numbers using -> function as an argument
	en2 := evenNumbers2([]int{12, 13, 13, 14, 15, 16, 17, 18, 19, 20},
		func(n []int) []int {
			var even []int
			for _, v := range n {
				if v%2 == 0 {
					even = append(even, v)
				}
			}
			return even
		})
	fmt.Println(en2)

	//? Exercise 2: filter a number from a list of numbers using higher order functions.
	filterNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})(22)
	filterNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})(10)
	// it will execute the closure function, with the value 22

}

///! important: Function Execution and Scope:
/*
	Normally, local variables (like sum in adder) would be destroyed once the function completes.
	However, because the returned anonymous function (closure) references sum, it keeps sum alive.
	The closure has a reference to the memory location where sum is stored.

	Why Captured Variables Remain Accessible?
	The key idea is that the returned anonymous function forms a closure, capturing the variables from its lexical scope. Here's what happens under the hood:

	Memory Allocation:
	When adder() is called, sum is allocated on the `[heap]` instead of the '[stack]' (because it needs to outlive the function call).
	The returned closure holds a reference to this heap-allocated sum.

	Variable Persistence:
	The closure holds a reference to sum's memory address. This means that sum's value is preserved across multiple calls to the closure.
	Even though adder() has finished executing, the memory for sum remains allocated because the closure still references it.
*/
