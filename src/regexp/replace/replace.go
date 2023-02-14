package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1, _ := regexp.Compile("\\.[a-zA-Z]+$")
	s1 := re1.ReplaceAllString("hello example.com", ".net")
	fmt.Println(s1)

	re2, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s2 := re2.ReplaceAllString("hello, world!", "${2}")
	fmt.Println(s2)

	re3, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s3 := re3.ReplaceAllString("hello, world!", "${2} ${1}")
	fmt.Println(s3)

	re4, _ := regexp.Compile("(?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)")
	s4 := re4.ReplaceAllString("hello, world!", "${second} ${first}")
	fmt.Println(s4)

	re5, _ := regexp.Compile("(?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)")
	s5 := re5.ReplaceAllLiteralString("hello, world!", "${second} ${first}")
	fmt.Println(s5)

	re6, _ := regexp.Compile("\\.([a-zA-Z]+)\\.")
	s6 := re6.Split("mail.hello.net www.example.com ftp.example.org", -1)
	fmt.Println(s6)

	re7, _ := regexp.Compile("\\.([a-zA-Z]+)\\.")
	s7 := re7.Split("mail.hello.net www.example.com, ftp.example.org", 2)
	fmt.Println(s7)

	re8, _ := regexp.Compile("\\.([a-zA-Z]+)\\.")
	s8 := re8.Split("mail.hello.net www.example.com, ftp.example.org", 3)
	fmt.Println(s8)

	re9, _ := regexp.Compile("\\.([a-zA-Z]+)\\.")
	s9 := re9.Split("mail.hello.net www.example.com, ftp.example.org", 0)
	fmt.Println(s9)
}
