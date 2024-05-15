package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)

}

// 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func a() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	// Output: 4 3 2 1 0
}
