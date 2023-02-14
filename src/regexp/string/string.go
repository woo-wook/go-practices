package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1, _ := regexp.Compile("\\.[a-zA-Z]+$")
	s1 := re1.FindString("hello example.com")
	fmt.Println(s1)

	re2, _ := regexp.Compile("([a-zA-Z]+)(\\.[a-zA-Z]+)$")
	s2 := re2.FindString("hello example.com")
	fmt.Println(s2)

	n2 := re2.FindStringSubmatchIndex("example.com")
	fmt.Println(n2)

	re3, _ := regexp.Compile("\\.[a-zA-Z.]+")
	s3 := re3.FindAllString("hello.co.kr world hello.net example.com", -1)
	fmt.Println(s3)

	s3 = re3.FindAllString("hello.co.kr world hello.net example.com", 2)
	fmt.Println(s3)

	n3 := re3.FindAllStringIndex("hello.co.kr world hello.net example.com", 2)
	fmt.Println(n3)
}
