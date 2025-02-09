package templates

import (
	"firstwails/domain"
	"html/template"
	"io/fs"
)

const modError = "ginserver:templates"

type templates struct {
	domain.IApp
	pages map[string]*template.Template
	fs    fs.FS
}

func New(app domain.IApp) *templates {
	t := &templates{
		IApp:  app,
		pages: make(map[string]*template.Template),
	}

	return t
}
