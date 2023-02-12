package main

import "fmt"

func main() {
	user := NewUser("test01", "Doyle")

	fmt.Println(user)
}

type User struct {
	id, name string
}

func NewUser(id string, name string) *User {
	return &User{id, name}
}
