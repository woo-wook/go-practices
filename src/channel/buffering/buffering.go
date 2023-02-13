package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2)
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("고루틴 :", i)
			time.Sleep(600)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		//time.Sleep(3000)
		fmt.Println("메인 함수 :", i)
	}

	fmt.Scanln()
}
