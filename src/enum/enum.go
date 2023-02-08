package main

import "fmt"

func main() {
	const (
		SUN = 0
		MON = 1
	)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Saturday)
	fmt.Println(numberOfDays)

	const (
		a = 1 << iota
		b
		c
		d
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	const (
		e, e1 = iota * 30, iota*30 + 1
		f, f1
		g, g1
		_, _
		i, i1
	)

	fmt.Println(e, e1)
	fmt.Println(f, f1)
	fmt.Println(g, g1)
	fmt.Println(i, i1)
}
