package main

import (
	"fmt"
	"os"
)

func main() {
	createFile()

	var num1 int
	var num2 float32
	var s string

	file1, _ := os.Open("src/filefunction/hello1.txt")
	defer file1.Close()
	n, _ := fmt.Fscan(file1, &num1, &num2, &s)

	fmt.Println("입력 개수:", n)
	fmt.Println(num1, num2, s)

	file2, _ := os.Open("src/filefunction/hello2.txt")
	defer file2.Close()
	n, _ = fmt.Fscanln(file2, &num1, &num2, &s)

	fmt.Println("입력 개수:", n)
	fmt.Println(num1, num2, s)

	file3, _ := os.Open("src/filefunction/hello3.txt")
	defer file3.Close()
	n, _ = fmt.Fscanln(file3, "%d,%f,%s", &num1, &num2, &s)

	fmt.Println("입력 개수:", n)
	fmt.Println(num1, num2, s)
}

func createFile() {
	file1, _ := os.Create("src/filefunction/hello1.txt")
	defer file1.Close()
	fmt.Fprint(file1, 1, 1.1, "Hello, World!")

	file2, _ := os.Create("src/filefunction/hello2.txt")
	defer file2.Close()
	fmt.Fprintln(file2, 1, 1.1, "Hello, World!")

	file3, _ := os.Create("src/filefunction/hello3.txt")
	defer file3.Close()
	fmt.Fprintf(file3, "%d,%f,%s", 1, 1.1, "Hello, World!")
}
