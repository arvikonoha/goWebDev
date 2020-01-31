package main

import (
	"fmt"
	"io"
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
		fmt.Fprint(conn, "Hello world")
		io.WriteString(conn, "\nBye world")
		conn.Close()
	}
}
