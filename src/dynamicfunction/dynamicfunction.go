package main

import (
	"fmt"
	"reflect"
)

func main() {
	var hello func()

	fn := reflect.ValueOf(&hello).Elem()

	v := reflect.MakeFunc(fn.Type(), h)

	fn.Set(v)

	hello()

	var intSum func(int, int) int64
	var stringSum func(string, string) string

	makeSum(&intSum)
	makeSum(&stringSum)

	fmt.Println(intSum(1, 2))
	fmt.Println(stringSum("Hello, ", "World!"))
}

func h(args []reflect.Value) []reflect.Value {
	fmt.Println("Hello, World!")
	return nil
}

func makeSum(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), sum)

	fn.Set(v)
}

func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]

	if a.Kind() != b.Kind() {
		fmt.Println("타입이 다르다.")
		return nil
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	}

	return []reflect.Value{}
}
