package main

import "fmt"

func main() {
	ageMap := map[string]int{"Doyle": 33, "Cary": 28}

	fmt.Println(ageMap)

	delete(ageMap, "Cary")

	fmt.Println(ageMap)
}
