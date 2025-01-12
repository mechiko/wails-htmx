package setup

import (
	"fmt"
	"html/template"
	"os"
)

func fileGetContents(filename string) string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("%s %s", err.Error(), filename))
	}
	return string(b)
}

var funcMapHtml = template.FuncMap{
	// The name "inc" is what the function will be called in the template text.
	"inc": func(i int) int {
		return i + 1
	},
	"noescape": func(str string) template.HTML {
		return template.HTML(str)
	},
}
