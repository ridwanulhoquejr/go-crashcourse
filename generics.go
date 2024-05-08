package main

// SumInits sums the values of a map of strings to integers.
func SumInits(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}
	return s
}

// ? Declare a type constraint
// isntead of declaring type constraints in the type parameter of a Generic function
// we can defince a interface of the types we want to use in the type parameter constraints
type Number interface {
	int64 | float64
}

// ? Generic func of above two functions
//
// [K comparable, V int64 | float64] this is called `[type parameter]` and
// comaparable and int64 | float64 are called `[type constraints]`
//
// Specify for the V type parameter a constraint that is a union of two types:
// int64 and float64. Using | specifies a union of the two types, meaning that this constraint allows either type. Either type will be permitted by the compiler as an argument in the calling code.
//
// if we defines what is inside type parameter then,
// we can use it in the function signature (i;e in the parameters)
func SumIntAndFloat[K comparable, V Number](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}
	return s
}

// if we define what is inside type parameter then,
// we can use it in the function signature (i;e in the parameters)
// in this case T is a type parameter and it is used in the parameters
func isLess[T Number](a, b T) bool {
	return a < b
}
