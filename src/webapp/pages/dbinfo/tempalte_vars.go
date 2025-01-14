package dbinfo

import (
	_ "embed"
)

//go:embed templates\index.html
var indexTmpl string

//go:embed templates\indexerror.html
var indexErrorTmpl string

//go:embed templates\error.html
var errorTmpl string

//go:embed templates\style.html
var styleTmpl string

//go:embed templates\navbar.html
var navbarTmpl string

//go:embed templates\page.html
var pageTmpl string

//go:embed templates\footer.html
var footerTmpl string
