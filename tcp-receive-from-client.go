package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	list, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
	}
	defer list.Close()
	for {
		conn, err := list.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
	defer conn.Close()
}
