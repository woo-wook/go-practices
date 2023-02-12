package main

import "fmt"

type Rectangle struct {
	width, height int
}

func main() {
	rectangle1 := Rectangle{30, 30}
	rectangle1.scaleA(10)
	fmt.Println(rectangle1)

	rectangle2 := Rectangle{30, 30}
	rectangle2.scaleB(10)
	fmt.Println(rectangle2)
}

func (rectangle Rectangle) scaleA(factor int) {
	rectangle.width = rectangle.width * factor
	rectangle.height = rectangle.height * factor
}

func (rectangle *Rectangle) scaleB(factor int) {
	rectangle.width = rectangle.width * factor
	rectangle.height = rectangle.height * factor
}
