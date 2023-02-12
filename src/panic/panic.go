package main

import "fmt"

func main() {
	f()

	fmt.Println("Hello!")
}

func f() {
	defer func() {
		s := recover()
		fmt.Println("[Recover]", s)
	}()

	a := [...]int{1, 2}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}
