// Package chunkio provides file splitting and joining functionality.
// It handles large files by dividing them into manageable chunks
// and reassembling them later.
package chunkio

// A file can be splitted into several chunks
// when the file is too large( more than 3GB in size)
// Each chunk represents a Gist entry in Github API
type Content struct {
	Chunks []Chunk
}

// Each Chunk can be splitted into several fragments
// when the chunk is too large(more than 10MB)
// Each fragment represents a file inside a single Gist
type Chunk struct {
	Fragment map[string]map[string]string
}
