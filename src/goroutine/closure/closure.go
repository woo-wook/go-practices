package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	s := "Hello, World!"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(s, i) // 전부 100!
		}()
	}

	fmt.Scanln()
}
