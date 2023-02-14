package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	s := "Hello, World!"
	reader := strings.NewReader(s)

	io.Copy(os.Stdout, reader)
}
