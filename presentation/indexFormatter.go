package presentation

import (
	"bytes"
	"html/template"
)

func index() string {
	tmpl := template.Must(template.ParseFiles("presentation/template/index.html"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, nil)
	return tpl.String()
}
