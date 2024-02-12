package main

type color map[string]string

func (c color) print() {
	for key, value := range c {
		println(key, value)
	}
}
