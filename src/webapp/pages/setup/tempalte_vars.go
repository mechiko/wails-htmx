package setup

import (
	_ "embed"
)

//go:embed templates\index.html
var indexTmpl string

//go:embed templates\page.html
var pageTmpl string

//go:embed templates\pagesave.html
var pageSaveTmpl string

//go:embed templates\modal.html
var modalTmpl string

//go:embed templates\omsid.html
var omsIdTmpl string

//go:embed templates\deviceid.html
var deviceIdTmpl string

//go:embed templates\configdb.html
var configDbTmpl string
