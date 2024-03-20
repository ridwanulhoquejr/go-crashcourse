package main

/// generic as a function
func sumSlice[T int | float32](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum

}

// gas-engine struc
type gasEngine struct {
	gallons float32
	mpg     float32
}

// electric-engine struc
type electricEngine struct {
	kwh   float32
	mpkwh float32
}

// now we can use generic to a specific property
type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T // this field `engine` now can use both T type [gasEngine or electricEngine] decided by the call side
}
