package main

import "fmt"

func main() {

	for bottles := 99; bottles >= 0; bottles-- {
		if bottles > 1 {
			fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", bottles, bottles)

			str := "bottles"

			if bottles-1 == 1 {
				str = "bottle"
			}

			fmt.Printf("Take one down, pass it around, %d %s of beer on the wall.\n", bottles, str)
		} else if bottles == 1 {
			fmt.Printf("%d bottle of beer on the wall, %d bottle of beer.\n", bottles, bottles)
			fmt.Printf("Take one down, pass it around, no more of beer on the wall.\n")
		} else {
			fmt.Printf("No more bottles of beer on the wall, no more bottles of beer.\n")
			fmt.Printf("Go to the store and bye some more, 99 bottles of beer on the wall.\n")
		}
	}
}
