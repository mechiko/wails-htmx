package embeded

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed root
var root embed.FS

func GetFileSystem() http.FileSystem {
	fsys, err := fs.Sub(root, "root")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
