package binfs

import (
	"bytes"
	"io"
	"os"
	"strings"
	"time"
)

type node struct {
	path  []string
	name  string
	nodes map[string]*node
	chunk *Chunk
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

func (n *node) ensureChild(name string) *node {
	if n.path == nil {
		n.path = []string{}
	}
	if n.nodes == nil {
		n.nodes = map[string]*node{}
	}
	c := n.nodes[name]
	if c == nil {
		c = &node{name: name, path: append(n.path, name)}
		n.nodes[name] = c
	}
	return c
}

func (n *node) ensure(name ...string) *node {
	t := n
	for _, v := range name {
		if v == "" {
			continue
		}
		t = t.ensureChild(v)
	}
	return t
}

func (n *node) find(name ...string) *node {
	t := n
	for _, v := range name {
		if v == "" {
			continue
		}
		if t != nil && t.nodes != nil {
			t = t.nodes[v]
		} else {
			return nil
		}
	}
	return t
}

func (n *node) toFileInfo() os.FileInfo {
	info := fileInfo{
		name: "/" + strings.Join(n.path, "/"),
	}
	if n.chunk != nil {
		info.date = n.chunk.Date
		info.size = int64(len(n.chunk.Data))
	} else {
		info.isDir = true
	}
	return info
}

func (n *node) toReadSeeker() io.ReadSeeker {
	if n.chunk != nil {
		return bytes.NewReader(n.chunk.Data)
	}
	return dirReadSeeker{}
}
