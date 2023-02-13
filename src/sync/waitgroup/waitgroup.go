package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	wg.Wait()
	fmt.Println("the end")
}
