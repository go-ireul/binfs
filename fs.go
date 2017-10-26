package binfs

import (
	"net/http"
)

type fileSystem struct{}

func (fileSystem) Open(name string) (http.File, error) {
	return Open(name)
}

// FileSystem creates http.FileSystem implementation
func FileSystem() http.FileSystem {
	return fileSystem{}
}
