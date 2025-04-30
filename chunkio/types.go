// Package chunkio provides file splitting and joining functionality.
// It handles large files by dividing them into manageable chunks
// and reassembling them later.
package chunkio

// A file can be splitted into several chunks(gist)
// if the file is too large( more than 3GB in size)
// Each chunk represents a Gist entry in Github API
type Content struct {
	Chunks []Chunk
}

// Each Chunk(Gist) can be splitted into several fragments(Files)
// when the chunk is too large(more than 10MB).
// A chunk can only contain a max number of 300 Fragments
// Making possible to store a 3GB file in single Gist
type Chunk struct {
	Fragments []Fragment
}

// Each fragment represents a file inside a single Gist
// and has a size limit of 10 MB
type Fragment struct {
	Number int
	Data   []byte
}
