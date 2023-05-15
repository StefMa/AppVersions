package presentation

import (
	"bytes"
	"html/template"
)

func index() string {
	tmpl := template.Must(template.ParseFS(templates, "template/index.html", "template/header.html"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, nil)
	return tpl.String()
}
