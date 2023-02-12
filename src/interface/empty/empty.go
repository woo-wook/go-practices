package main

import (
	"fmt"
	"strconv"
)

type Any interface {
}

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello"))
	fmt.Println(formatString(Person{"Doyle", 30}))
	fmt.Println(formatString(Person{"Cary", 28}))
}

func formatString(arg Any) string {
	switch arg.(type) {
	case int:
		i := arg.(int)
		return strconv.Itoa(i)
	case float32:
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 32)
	case string:
		s := arg.(string)
		return s
	case Person:
		p := arg.(Person)
		return p.name + " " + strconv.Itoa(p.age)
	case *Person:
		p := arg.(*Person)
		return p.name + " " + strconv.Itoa(p.age)
	default:
		return "Error"
	}
}
