package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	myfile, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}
	defer myfile.Close()
	message := `<html><body><h1>` + os.Args[1] + `</h1></body></html>`
	io.Copy(myfile, strings.NewReader(message))
}
