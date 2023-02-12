package main

import "fmt"

func main() {
	rectangle := Rectangle{30, 30}
	area := rectangleArea(&rectangle)

	fmt.Println(area)
}

type Rectangle struct {
	width, height int
}

func rectangleArea(rect *Rectangle) int {
	return rect.width * rect.height
}
