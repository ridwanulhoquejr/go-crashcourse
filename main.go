package main

import "fmt"

func main() {

	//  initialize a map of strings to integers
	mapInt := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// initialize a map of strings to floats
	mapFloat := map[string]float64{
		"one1":  1.1,
		"two":   2.2,
		"three": 3.3,
	}

	fmt.Println(SumInits(mapInt))    // 6
	fmt.Println(SumFloats(mapFloat)) // 6.6

	// Generic func of above two functions
	fmt.Println(SumIntAndFloat[string, int64](mapInt))     // 6
	fmt.Println(SumIntAndFloat[string, float64](mapFloat)) // 6.6

	// isLess func
	fmt.Println(isLess[int64](1, 2))       // true
	fmt.Println(isLess[float64](1.5, 1.7)) // true
}
