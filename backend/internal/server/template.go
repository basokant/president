package server

import (
	_ "embed"
	"text/template"
)

//go:embed "templates/index.html"
var html_page string

var html = template.Must(template.New("chat_room").Parse(html_page))
