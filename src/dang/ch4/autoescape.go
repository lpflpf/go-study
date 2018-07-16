// autoescape
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const temp1 = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`

	t := template.Must(template.New("escape").Parse(temp1))
	var data struct {
		A string
		B template.HTML // 可以直接使用，不会做转义
	}

	data.A = "<b>Hello !</b>"
	data.B = "<b>Hello!</b"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
