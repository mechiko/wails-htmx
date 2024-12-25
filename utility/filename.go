package utility

import (
	"path/filepath"
	"regexp"
	"strings"
)

func ClearForFileName(str string) string {
	str = regexp.MustCompile(`[^\p{L}\p{N} ]+`).ReplaceAllString(str, "")
	return str
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
}
