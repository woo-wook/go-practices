package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := []string{"Maria", "Doyle", "Cary", "James"}

	r := ring.New(len(data))

	for i := 0; i < r.Len(); i++ {
		r.Value = data[i]
		r = r.Next()
	}

	r.Do(func(x any) {
		fmt.Println(x)
	})

	fmt.Println("Move forward : ")
	r = r.Move(1)

	fmt.Println("Curr :", r.Value)
	fmt.Println("Prev :", r.Prev().Value)
	fmt.Println("Next :", r.Next().Value)
}
