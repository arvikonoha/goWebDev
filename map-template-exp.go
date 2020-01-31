package main

import (
	"fmt"
	"os"
	"text/template"
)

var tlp *template.Template

func init() {
	tlp = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	err := tlp.ExecuteTemplate(os.Stdout, "mapTemplate.gohtml", map[string]string{
		"hello":   "world",
		"goodbye": "faggot",
		"adios":   "amigos",
	})

	if err != nil {
		fmt.Println(err)
	}
}
