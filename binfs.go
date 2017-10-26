package binfs

import (
	"time"
)

// Chunk a file in a binfs
type Chunk struct {
	Path []string
	Date time.Time
	Data []byte
}

// DefaultRoot the default root of BinFS
var DefaultRoot = &Node{}

// Load load a file into zone
func Load(c *Chunk) {
	DefaultRoot.Load(c)
}

// Open open a file, a partial mocking of *os.File
func Open(name string) (File, error) {
	return DefaultRoot.Open(name)
}

// Find find a deep child node
func Find(name ...string) *Node {
	return DefaultRoot.Find(name...)
}

// Walk walk the default root
func Walk(fn NodeWalker) {
	DefaultRoot.Walk(fn)
}
