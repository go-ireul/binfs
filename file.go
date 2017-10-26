package binfs

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// ErrIsDirectory error returned while trying read/seek a directory
var ErrIsDirectory = errors.New("is a directory")

// File abstracts a binfs file
type File interface {
	http.File
}

// dummy io.ReadSeeker for directory
type dirReadSeeker struct{}

func (dirReadSeeker) Read(p []byte) (n int, err error) {
	return 0, ErrIsDirectory
}

func (dirReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, ErrIsDirectory
}

// fileInfo implements os.FileInfo
type fileInfo struct {
	name  string
	size  int64
	date  time.Time
	isDir bool
}

func (f fileInfo) Name() string {
	return f.name
}

func (f fileInfo) Size() int64 {
	return f.size
}

func (f fileInfo) Mode() os.FileMode {
	if f.isDir {
		return os.FileMode(0777)
	}
	return os.FileMode(0666)
}

func (f fileInfo) ModTime() time.Time {
	return f.date
}

func (f fileInfo) IsDir() bool {
	return f.isDir
}

func (fileInfo) Sys() interface{} {
	return nil
}

type file struct {
	io.ReadSeeker
	path []string
	info fileInfo
}

// Close close implements io.Closer
func (f file) Close() error {
	return nil
}

func (f file) Readdir(n int) ([]os.FileInfo, error) {
	return []os.FileInfo{}, nil
}

func (f file) Stat() (os.FileInfo, error) {
	return f.info, nil
}

// newFile creates a file from a node
func newFile(n *node) (*file, error) {
	f := &file{
		path: n.path,
		info: fileInfo{
			name: "/" + strings.Join(n.path, "/"),
		},
	}
	if n.chunk != nil {
		f.ReadSeeker = bytes.NewReader(n.chunk.Data)
		f.info.date = n.chunk.Date
		f.info.size = int64(len(n.chunk.Data))
	} else {
		f.ReadSeeker = dirReadSeeker{}
		f.info.isDir = true
	}
	return f, nil
}
