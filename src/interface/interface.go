package main

import "fmt"

type MyInt int // 기존 자료형을 새 자료형으로 선언하기

type Rectangle struct {
	width, height int
}

type hello interface {
}

type Printer interface {
	Print()
}

func main() {
	var h hello
	fmt.Println(h)

	var i MyInt = 5
	r := Rectangle{10, 30}

	var p Printer

	p = i
	p.Print()

	p = r
	p.Print()

	printer1 := Printer(i)
	printer2 := Printer(r)

	printer1.Print()
	printer2.Print()
}

func (i MyInt) Print() {
	fmt.Println(i)
}

func (i Rectangle) Print() {
	fmt.Println(i.width, i.height)
}
