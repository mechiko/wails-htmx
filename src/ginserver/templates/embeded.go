package templates

import (
	"embed"
)

//go:embed "dbinfo" "finder" "setup" "stats" "index"
var root embed.FS
