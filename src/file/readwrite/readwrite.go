package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("src/file/file.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(644))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	n := 0
	s := "Hello"

	n, err = file.Write([]byte(s))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")

	fi, err := file.Stat()

	if err != nil {
		fmt.Println(err)
		return
	}

	data := make([]byte, fi.Size())

	file.Seek(0, io.SeekStart)

	n, err = file.Read(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data))
}
