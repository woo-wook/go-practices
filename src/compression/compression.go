package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	err := makeFile()

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := readFile()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}

func readFile() (string, error) {
	file, err := os.Open("src/compression/hello.txt.gz")

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer file.Close()

	reader, err := gzip.NewReader(file)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer reader.Close()

	read, err := io.ReadAll(reader)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(read), nil
}

func makeFile() error {
	file, err := os.OpenFile("src/compression/hello.txt.gz", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))

	if err != nil {
		return err
	}

	defer file.Close()

	writer := gzip.NewWriter(file)

	defer writer.Close()

	s := "Hello, World!"
	writer.Write([]byte(s))
	writer.Flush()

	return nil
}
