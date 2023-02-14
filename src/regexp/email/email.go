package main

import (
	"fmt"
	"regexp"
)

func main() {
	var validEmail, _ = regexp.Compile(
		"^[_a-z0-9+-.]+@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$",
	)

	fmt.Println(validEmail.MatchString("hello@mail.example.com"))
	fmt.Println(validEmail.MatchString("hello+kr@mail.example.com"))
	fmt.Println(validEmail.MatchString("hello-world@mail.example.com"))
	fmt.Println(validEmail.MatchString("hello_1@mail.example.com"))
	fmt.Println(validEmail.MatchString("hello_1@e-xample.co"))

	fmt.Println(validEmail.MatchString("@e-xample.co"))
	fmt.Println(validEmail.MatchString("hello@e-xample."))
	fmt.Println(validEmail.MatchString("hello@e-xample.aaaaa"))
}
