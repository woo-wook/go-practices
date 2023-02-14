package main

import (
	"fmt"
	"strconv"
)

func main() {
	var err error

	var num1 int
	num1, err = strconv.Atoi("100")
	fmt.Println(num1, err)

	var num2 int
	num2, err = strconv.Atoi("10t")
	fmt.Println(num2, err)

	var s1 string
	s1 = strconv.Itoa(1000)
	fmt.Println(s1)

	var s2 string
	s2 = strconv.FormatBool(true)
	fmt.Println(s2)

	var s3 string
	s3 = strconv.FormatInt(-10, 10)
	fmt.Println(s3)

	var s4 string
	s4 = strconv.FormatUint(32, 16)
	fmt.Println(s4)

	var s = make([]byte, 0)

	s = strconv.AppendBool(s, true)
	fmt.Println(string(s))

	s = strconv.AppendFloat(s, 1.3, 'f', -1, 32)
	fmt.Println(string(s))

	s = strconv.AppendInt(s, -10, 10)
	fmt.Println(string(s))

	s = strconv.AppendUint(s, 32, 16)
	fmt.Println(string(s))

	var b1 bool
	b1, err = strconv.ParseBool("true")
	fmt.Println(b1, err)

	var num4 float64
	num4, err = strconv.ParseFloat("1.3", 64)
	fmt.Println(num4, err)

	var num5 int64
	num5, err = strconv.ParseInt("-10", 10, 32)
	fmt.Println(num5, err)

	var num6 uint64
	num6, err = strconv.ParseUint("20", 16, 32)
	fmt.Println(num6, err)
}
