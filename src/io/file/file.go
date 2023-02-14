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

	writer := bufio.NewWriter(file)

	writer.WriteString("Hello, World!")
	writer.Flush()

	reader := bufio.NewReader(file)

	fi, _ := file.Stat()
	b := make([]byte, fi.Size())

	file.Seek(0, io.SeekStart)
	reader.Read(b)

	fmt.Println(string(b))
}
