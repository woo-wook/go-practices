package main

import (
	"fmt"
	"sync"
)

func main() {
	once := new(sync.Once)

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("goroutine :", n)

			once.Do(hello)
		}(i)
	}

	fmt.Scanln()
}

func hello() {
	fmt.Println("Hello, World!")
}
