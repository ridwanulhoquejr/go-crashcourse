package main

import "fmt"

// Question:
// Filter a number from a list of numbers using higher order functions.

func filterNumber(n []int) func(x int) {
	var found bool
	return func(x int) {
		for _, v := range n {
			if v == x {
				found = true
				// break
			} else {
				found = false
				// break
			}
		}

		if found {
			fmt.Println("Number Not Found: ", x)
		} else {
			fmt.Println("Number Found: ", x)
		}
	}
}
