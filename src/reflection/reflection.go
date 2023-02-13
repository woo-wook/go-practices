package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a, b int
}

func main() {
	num := 1
	fmt.Println(reflect.TypeOf(num))

	s := "Hello, World!"
	fmt.Println(reflect.TypeOf(s))

	var f float32 = 1.3
	fmt.Println(reflect.TypeOf(f))

	data := Data{1, 2}
	fmt.Println(reflect.TypeOf(data))

	float2 := 1.3

	t := reflect.TypeOf(float2)
	v := reflect.ValueOf(float2)

	fmt.Println(t.Name())
	fmt.Println(t.Size())
	fmt.Println(t.Kind() == reflect.Float64)
	fmt.Println(t.Kind() == reflect.Int64)

	fmt.Println(v.Type())
	fmt.Println(v.Kind() == reflect.Float64)
	fmt.Println(v.Kind() == reflect.Int64)
	fmt.Println(v.Float())
}
