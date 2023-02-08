package main

import (
	"fmt"
	"strconv"
)

func main() {
	const age int = 10
	const name string = "Maria"
	//const score int // 컴파일 오류

	const address = "서울시 구로구.."

	const x, y int = 30, 50

	const (
		j, i     = 30, 20
		user, no = "Doyle", 300
	)

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(x + y)
	fmt.Println(j + i)
	fmt.Println(user + strconv.Itoa(no))
}
