package binfs

import (
	"errors"
	"io"
	"net/http"
	"os"
)

// ErrIsDirectory error returned while trying read/seek a directory
var ErrIsDirectory = errors.New("is a directory")

// File abstracts a binfs file
type File interface {
	http.File
}

type file struct {
	io.ReadSeeker
	info os.FileInfo
	path []string
}

// Close close implements io.Closer
func (f file) Close() error {
	return nil
}

func (f file) Readdir(n int) ([]os.FileInfo, error) {
	c := fsRoot.find(f.path...)
	if n > 0 {
		return []os.FileInfo{}, errors.New("n > 0 is not supported")
	}
	// returns EOF if not found or n is too large
	if c == nil || c.nodes == nil {
		return []os.FileInfo{}, nil
	}
	// dir nodes
	out := []os.FileInfo{}
	for _, sub := range c.nodes {
		out = append(out, sub.toFileInfo())
	}
	return out, nil
}

func (f file) Stat() (os.FileInfo, error) {
	return f.info, nil
}

// newFile creates a file from a node
func newFile(n *node) *file {
	return &file{
		ReadSeeker: n.toReadSeeker(),
		path:       n.path,
		info:       n.toFileInfo(),
	}
}
