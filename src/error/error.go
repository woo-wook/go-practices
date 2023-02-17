package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	s, err := helloOne(1)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("s: ", s)

	s, err = helloOne(2)

	if err != nil {
		log.Panicln("err", err)
	}

	// 에러로 종료되어서 실행 X
	fmt.Println(s)
	fmt.Println("HelloWorld!")
}

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}

	return "", fmt.Errorf("%d는 1이 아닙니다\n", n)
}
