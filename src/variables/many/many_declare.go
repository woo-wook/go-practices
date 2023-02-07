package main

import "fmt"

func main() {
	var x, y int = 30, 50
	var age, name = 10, "Doyle"

	a, b, c, d := 1, 3.4, "Hello", false

	fmt.Printf("x = %d, y = %d, age = %d, name = %s, a = %d, b = %f, c = %s, d = %d\n", y, x, age, name, a, b, c, d)

	x, y, age = 30, 50, 5

	fmt.Printf("x = %d, y = %d, age = %d, name = %s, a = %d, b = %f, c = %s, d = %d\n", y, x, age, name, a, b, c, d)

	var (
		j, i     int = 10, 20
		user, no     = "Doyle", 300
	)

	fmt.Printf("j = %d, i = %d, user = %s, no = %d", j, i, user, no)

	z := 300
	_ = z
}
