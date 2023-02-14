package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("src/io/file.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	s := "Hello, World!"
	r := strings.NewReader(s)

	w := bufio.NewWriter(file)

	w.ReadFrom(r)
	w.Flush()

	// 방법 2
	io.Copy(w, r)
}
