package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8000") // 포트만 입력하면 모든 데이터를 수신하고, 아이피를 입력하면 해당 아이피에서만 수신한다.

	if err != nil {
		fmt.Println(err)
		return
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		defer conn.Close()

		go requestHandler(conn)
	}
}

func requestHandler(c net.Conn) {
	data := make([]byte, 4096)

	for {
		n, err := c.Read(data)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n]))

		_, err = c.Write(data[:n])

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
