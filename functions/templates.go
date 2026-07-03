package functions

import "html/template"

var tmpl = template.Must(
	template.ParseFiles("templates/index.html"),
)
