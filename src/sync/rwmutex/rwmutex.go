package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	race()
	rwmutex()
}

func rwmutex() {
	data := 0
	var mutex sync.RWMutex

	go func() {
		for i := 0; i < 3; i++ {
			mutex.Lock()
			data += 1
			fmt.Println("write :", data)
			time.Sleep(10 * time.Millisecond)
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			mutex.RLock()
			fmt.Println("read 1 :", data)
			time.Sleep(1 * time.Millisecond)
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			mutex.RLock()
			fmt.Println("read 2 :", data)
			time.Sleep(2 * time.Millisecond)
			mutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
}

func race() {
	data := 0

	go func() {
		for i := 0; i < 3; i++ {
			data += 1
			fmt.Println("write :", data)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("read 1 :", data)
			time.Sleep(1 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("read 2 :", data)
			time.Sleep(2 * time.Millisecond)
		}
	}()

	time.Sleep(10 * time.Second)
}
