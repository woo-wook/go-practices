package main

import "fmt"

func main() {
	i := 10

	if i >= 5 {
		fmt.Println("5 이상입니다.")
	} else if i >= 10 {
		fmt.Println("10 이상입니다.")
	} else {
		fmt.Println("모르겠습니다.")
	}
}
