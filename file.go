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
	// idx defines current cursor while executing Readdir(n int)
	idx int
}

// Close close implements io.Closer
func (f file) Close() error {
	return nil
}

func (f file) Readdir(n int) ([]os.FileInfo, error) {
	// out
	out := []os.FileInfo{}
	// find nodes
	nodes := fsRoot.find(f.path...).sortedNodes()
	// handle n > 0
	if n > 0 {
		var err error
		// if empty
		if len(nodes) == 0 {
			return out, io.EOF
		}
		// determine iteration max
		max := f.idx + n
		if max > len(nodes) {
			max = len(nodes)
			err = io.EOF
		}
		// output
		for i := f.idx; i < max; i++ {
			out = append(out, nodes[i].toFileInfo())
		}
		f.idx = max - 1
		return out, err
	}
	// dir all nodes
	for _, sub := range nodes {
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
