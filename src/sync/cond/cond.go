package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex = new(sync.Mutex)
	cond := sync.NewCond(mutex)

	c := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()
			c <- true
			fmt.Println("wait begin :", n)
			cond.Wait()
			fmt.Println("wait end :", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c
	}
	//
	//for i := 0; i < 3; i++ {
	//	mutex.Lock()
	//	fmt.Println("signal :", i)
	//	cond.Signal()
	//	mutex.Unlock()
	//}

	mutex.Lock()
	fmt.Println("broadcast")
	cond.Broadcast()
	mutex.Unlock()

	fmt.Scanln()
}
