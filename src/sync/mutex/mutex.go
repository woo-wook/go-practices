package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	race()
	usingMutex()
}

func usingMutex() {
	var data []int
	mutex := new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()

			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()

			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(len(data))
}

func race() {
	var data []int

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)

			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)

			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(len(data))
}
