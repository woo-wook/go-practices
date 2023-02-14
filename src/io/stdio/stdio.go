package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "Hello, World!"
	reader := strings.NewReader(s)

	var s1, s2 string
	n, _ := fmt.Fscanf(reader, "%s %s", &s1, &s2)

	fmt.Println("입력 수:", n)
	fmt.Println(s1)
	fmt.Println(s2)

	file, err := os.OpenFile("src/io/file.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintf(writer, "%s %s", s1, s2)
	writer.Flush()
}
