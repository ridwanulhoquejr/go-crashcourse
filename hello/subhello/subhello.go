package subhello

import "fmt"

func FromSubHello(s string) string {

	m := fmt.Sprintf("Hello, %v. i am from subhello", s)
	return m
}
