package main

// pass by references
type Str struct {
	number int
}

func addOneWithParams(x *int) {
	*x++
}

func (p *Str) addOneWithStruct() {
	p.number++
}
