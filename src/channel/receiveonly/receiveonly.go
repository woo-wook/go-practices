package main

import "fmt"

func main() {
	channel1 := make(chan int)

	go producer(channel1)
	go customer(channel1)

	fmt.Scanln()

	channel2 := sum(1, 2)

	fmt.Println(<-channel2)
	fmt.Scanln()

	channel3 := num(1, 2)
	out := sumToChannel(channel3)

	fmt.Println(<-out)
}

func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100

	// fmt.Println(<-c) // 컴파일 에러 발생
}

func customer(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c)

	//c <- 1 // 컴파일 에러 발생
}

func sum(a, b int) <-chan int {
	out := make(chan int)

	go func() {
		out <- a + b
	}()

	return out
}

func num(a, b int) <-chan int {
	out := make(chan int)

	go func() {
		out <- a
		out <- b
		close(out)
	}()

	return out
}

func sumToChannel(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		r := 0
		for i := range c {
			r = r + i
		}

		out <- r
	}()

	return out
}
