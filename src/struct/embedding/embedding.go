package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	school string
	grade  int
}

func main() {
	var s = Student{grade: 3, school: "S", Person: Person{name: "Doyle", age: 30}}

	fmt.Println(s)

	s.Person.greeting()
	s.greeting()
}

func (_ *Person) greeting() {
	fmt.Println("Hello!")
}

func (_ *Student) greeting() {
	fmt.Println("Hello Student!")
}
