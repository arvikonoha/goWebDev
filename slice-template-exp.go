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
	err := tlp.ExecuteTemplate(os.Stdout, "sliceTemplate.gohtml", []string{"hello", "world"})
	if err != nil {
		fmt.Println(err)
	}
}
