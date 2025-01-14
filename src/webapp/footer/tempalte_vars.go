package footer

import (
	_ "embed"
)

//go:embed templates\content.html
var contentTmpl string

//go:embed templates\modal.html
var modalTmpl string
