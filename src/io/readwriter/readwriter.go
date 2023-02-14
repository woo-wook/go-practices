package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("src/io/file.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(file)

	readWriter := bufio.NewReadWriter(reader, writer)

	readWriter.WriteString("Hello, World!")
	readWriter.Flush()

	file.Seek(0, io.SeekStart)

	data, _, _ := readWriter.ReadLine()

	fmt.Println(string(data))
}
