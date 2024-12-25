package home

import (
	_ "embed"
)

//go:embed templates\index.html
var indexTmpl string

//go:embed templates\style.html
var styleTmpl string
