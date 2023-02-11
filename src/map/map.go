package main

import "fmt"

func main() {
	var m1 map[string]int32
	m2 := make(map[string]int32)
	m3 := map[string]int{
		"Ryu":   32,
		"Doyle": 55,
	}

	fmt.Println(m3["Ryu"])
	fmt.Println(m3["Doyle"])

	_, _ = m1, m2
}
