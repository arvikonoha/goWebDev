package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var fn = template.FuncMap{
	"uc": strings.ToUpper,
}

var tlp *template.Template

func init() {
	tlp = template.Must(template.New("").Funcs(fn).ParseGlob("./templates/*"))
}

func main() {
	err := tlp.ExecuteTemplate(os.Stdout, "funcTemplete.gohtml", "hello")
	if err != nil {
		fmt.Println(err)
	}
}
