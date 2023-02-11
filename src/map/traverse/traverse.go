package main

import "fmt"

func main() {

	ageMap := map[string]int{"Doyle": 33, "Cary": 28}

	for name, age := range ageMap {
		fmt.Println(name, age)
	}
}
