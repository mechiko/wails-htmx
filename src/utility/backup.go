package utility

import (
	"fmt"
	"time"
	_ "time/tzdata"
)

func TimeFileName(name string) string {
	n := time.Now().Local()
	return fmt.Sprintf("%s_%4d.%02d.%02d_%02d%02d%02d", name, n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Second())
}
