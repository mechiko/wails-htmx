package scrolling

import (
	_ "embed"
)

//go:embed views\blocks-index.html
var blocksIndexTmpl string

//go:embed views\blocks.html
var blocksTmpl string

//go:embed views\style.html
var styleTmpl string

//go:embed views\input-query.html
var inputQueryTmpl string

//go:embed views\select-columns.html
var selectColumns string
