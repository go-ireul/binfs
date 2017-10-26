package binfs

import (
	"io"
	"net/http"
	"os"
)

// File abstracts a binfs file
type File interface {
	http.File
}

type file struct {
	io.Reader
	io.Seeker
	name  string
	isDir bool
}

// Close close implements io.Closer
func (f *file) Close() error {
	return nil
}

func (f *file) Readdir(n int) ([]os.FileInfo, error) {
	return []os.FileInfo{}, nil
}

func (f *file) Stat() (os.FileInfo, error) {
	return nil, nil
}
