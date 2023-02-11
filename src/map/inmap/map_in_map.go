package main

import "fmt"

func main() {
	userMap := map[int]map[string]string{
		1: {
			"name":   "Doyle",
			"userId": "doyle01",
		},
		2: {
			"name":   "Cary",
			"userId": "Cary01",
		},
	}

	fmt.Println(userMap)

	for key, value := range userMap {
		fmt.Println("UserNo", key)

		for property, propertyValue := range value {
			fmt.Printf("%s: %s\n", property, propertyValue)
		}
	}

	fmt.Println(userMap[1]["name"])
}
