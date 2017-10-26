package binfs

import (
	"net/http"
)

type fs struct{}

func (fs) Open(name string) (http.File, error) {
	return Open(name)
}

// FileSystem creates http.FileSystem implementation
func FileSystem() http.FileSystem {
	return fs{}
}
