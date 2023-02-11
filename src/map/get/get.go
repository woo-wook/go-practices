package main

import "fmt"

func main() {
	empMap := make(map[int32]string)

	empMap[3000] = "Doyle"
	empMap[5000] = "Cary"
	empMap[7000] = "Mary"

	fmt.Println(empMap[3000])
	fmt.Println(empMap[5000])
	fmt.Println(empMap[7000])

	if value, hasKey := empMap[7000]; hasKey {
		fmt.Println(value)
	}
}
