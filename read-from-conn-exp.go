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
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println(msg)
	}
	defer conn.Close()
}
