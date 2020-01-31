package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Args[0])
	} else {
		fmt.Printf("%T", os.Args[1])
	}
}
