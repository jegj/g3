// Package chunkio provides file splitting and joining functionality.
// It handles large files by dividing them into manageable chunks
// and reassembling them later.
package chunkio

// Content can be splitted into several fragments
// Depending of the Fragments length( 300 files of 10MB), we can redefine
// this structure as several Gists
type Content struct {
	Fragments []Fragment
}

// Fragment represents a file inside a single Gist
// and has a size limit of 10 MB
type Fragment struct {
	Number   int
	Filepath string
}

type ChunkProcessor func(chunk []byte, chunkIndex int) error
