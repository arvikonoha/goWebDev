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
	err := tlp.ExecuteTemplate(os.Stdout, "structTemplate.gohtml", struct {
		Name string
		Age  int
	}{Name: "Aravind", Age: 10})
	if err != nil {
		fmt.Println(err)
	}
}
