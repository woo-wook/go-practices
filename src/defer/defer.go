package main

import (
	"fmt"
	"os"
)

func main() {
	defer world()
	defer hello()

	readHello()
}

func readHello() {
	file, err := os.Open("src/defer/hello.txt")

	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 100)

	if _, err := file.Read(buffer); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer))
}

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}
