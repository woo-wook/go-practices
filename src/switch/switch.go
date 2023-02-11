package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	s := "Hello"

	switch s {
	case "Hello":
		fmt.Println("Hello")
		fallthrough
	case "World":
		fmt.Println("World")
	default:
		fmt.Println("default")
	}

	num := 30

	switch num {
	case 30, 40, 50:
		fmt.Println("30, 40, 50")
	case 60, 70:
		fmt.Println("60. 70")
	}

	switch {
	case num >= 5 && num < 50:
		fmt.Println("5 이상, 50 미만")
	case num >= 0 && num < 5:
		fmt.Println("0 이상, 5 미만")
	}

	rand.Seed(time.Now().UnixNano())

	switch i := rand.Intn(10); {
	case i >= 3 && i < 6:
		fmt.Println("3 이상, 6 미만")
	case i == 9:
		fmt.Println("9")
	default:
		fmt.Println(i)
	}
}
