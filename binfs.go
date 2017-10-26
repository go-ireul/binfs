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

// Load load a file into zone
func Load(f *Chunk) {
}

// Open open a file, a partial mocking of *os.File
func Open(name string) (File, error) {
	return &file{}, nil
}
