package main

import "fmt"

func main() {
	var s1, s2 string
	n, _ := fmt.Scan(&s1, &s2)

	fmt.Println("입력 개수: ", n)
	fmt.Println(s1, s2)

	var s3, s4 string
	c, _ := fmt.Scanln(&s3, &s4)

	fmt.Println("입력 개수: ", c)
	fmt.Println(s3, s4)

	var num1, num2 int
	p, _ := fmt.Scanf("%d,%d", &num1, &num2)

	fmt.Println("입력 개수: ", p)
	fmt.Println(num1, num2)
}
