package main

import "fmt"

func main() {
	var s1 string
	s1 = fmt.Sprint(1, 1.1, "Hello, World!")
	fmt.Print(s1)

	var s2 string
	s2 = fmt.Sprintln(1, 1.1, "Hello, World!")
	fmt.Print(s2)

	var s3 string
	s3 = fmt.Sprintf("%d %f %s\n", 1, 1.1, "Hello, World!")
	fmt.Print(s3)

	var num1 int
	var num2 float32
	var s string

	input1 := "1\n1.1\nhello"
	n, _ := fmt.Sscan(input1, &num1, &num2, &s)

	fmt.Println("입력 개수 :", n)
	fmt.Println(num1, num2, s)

	input2 := "1 1.1 hello"
	n, _ = fmt.Sscanln(input2, &num1, &num2, &s)

	fmt.Println("입력 개수 :", n)
	fmt.Println(num1, num2, s)

	input3 := "1,1.1,hello"
	n, _ = fmt.Sscanln(input3, "%d,%f,%s", &num1, &num2, &s)

	fmt.Println("입력 개수 :", n)
	fmt.Println(num1, num2, s)
}
