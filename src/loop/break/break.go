package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 4 {
			break
		}

		fmt.Println(i)
		i++
	}

loop:
	for i := 0; i < 3; i++ {
	loop2:
		for j := 0; j < 3; j++ {
			if j == 2 {
				break loop
			}

			if i == 2 {
				break loop2
			}

			fmt.Println(i, j)
		}
	}
}
