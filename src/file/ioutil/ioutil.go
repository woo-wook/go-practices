package main

import (
	"fmt"
	"os"
)

func main() {
	s := "Hello, World!"

	err := os.WriteFile("src/file/file.txt", []byte(s), os.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := os.ReadFile("src/file/file.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
