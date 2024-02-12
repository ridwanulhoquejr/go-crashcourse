package main

import "fmt"

type logWriter struct{}

// now, logWriter is implementing the Writer interface
func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
