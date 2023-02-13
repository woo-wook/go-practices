package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"Name"`
	age  int    `tag1:"나이" tag2:"age"`
}

func main() {
	person := Person{}

	name, ok := reflect.TypeOf(person).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	name, ok = reflect.TypeOf(person).FieldByName("age")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))
}
