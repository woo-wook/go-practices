package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{10, 6, 3, 7, 9}
	b := []float64{4.2, 30.1, 22.3, 3315.1, 0.3}
	c := []string{"Maria", "Doyle", "Cary"}

	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)

	sort.Sort(sort.Float64Slice(b))
	fmt.Println(b)
	sort.Sort(sort.Reverse(sort.Float64Slice(b)))
	fmt.Println(b)

	sort.Sort(sort.StringSlice(c))
	fmt.Println(c)
	sort.Sort(sort.Reverse(sort.StringSlice(c)))
	fmt.Println(c)
}
