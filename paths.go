package baghl

import (
	"path/filepath"
	"strings"
)

func FilesystemRouteToURLEndpoint(path string) string {
	newPath := filepath.ToSlash(path)
	newPath = strings.Replace(newPath, "routes/", "", 1)
	return newPath
}
