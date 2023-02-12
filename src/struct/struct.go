package main

import "fmt"

type Rectangle struct {
	width, height int
}

func main() {
	var rect Rectangle
	var rectPtr *Rectangle = new(Rectangle)

	fmt.Println(rect)
	fmt.Println(rectPtr)

	var rect1 = Rectangle{10, 20}
	rect2 := Rectangle{width: 50, height: 30}

	fmt.Println(rect1)
	fmt.Println(rect2)
}
