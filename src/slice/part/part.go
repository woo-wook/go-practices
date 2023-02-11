package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a[0:3])
	fmt.Println(a[0:5])
	fmt.Println(a[1:5])
	fmt.Println(a[:])
	fmt.Println(a[0:])
	fmt.Println(a[:3])
	fmt.Println(a[:len(a)])
	fmt.Println(a[0:5:cap(a)])
	fmt.Println(a[:5:cap(a)])
}
