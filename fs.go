package binfs

import (
	"net/http"
)

type httpFileSystem struct{}

func (httpFileSystem) Open(name string) (http.File, error) {
	return Open(name)
}

// FileSystem creates http.FileSystem implementation
func FileSystem() http.FileSystem {
	return httpFileSystem{}
}
