package binfs

import (
	"os"
	"strings"
	"time"
)

// Chunk a file in a binfs
type Chunk struct {
	Path []string
	Date time.Time
	Data []byte
}

var fsRoot = &node{}

// Load load a file into zone
func Load(c *Chunk) {
	fsRoot.ensure(c.Path...).chunk = c
}

// Open open a file, a partial mocking of *os.File
func Open(name string) (File, error) {
	comps := strings.Split(name, "/")
	n := fsRoot.find(comps...)
	if n == nil {
		return nil, os.ErrNotExist
	}
	return newFile(n), nil
}
