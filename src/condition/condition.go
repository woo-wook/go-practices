package main

import (
	"fmt"
	"os"
)

func main() {
	i := 10

	if i >= 5 {
		fmt.Println("5 이상입니다.")
	} else if i >= 10 {
		fmt.Println("10 이상입니다.")
	} else {
		fmt.Println("모르겠습니다.")
	}

	if file, err := os.ReadFile("src/condition/hello.txt"); err == nil {
		fmt.Printf("%s", file)
	}
}
