package chunkio

import (
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	ChunkSize   = 10 * 1024 * 1024 // 10MB chunks
	BufferSize  = 5 * 1024 * 1024  // 5MB buffer
	WorkerCount = 4                // Number of concurrent workers
)

var (
	globalContent = Content{}
	contentMutex  sync.RWMutex
)

// TODO: Check whyt number chunks does not match file size
func SplitFileConcurrent(inputPath string) (*Content, error) {
	fmt.Println("Splitting file", inputPath)
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	// Worker pool pattern
	chunkChan := make(chan Fragment, WorkerCount)
	errChan := make(chan error, 1)
	doneChan := make(chan struct{})
	var wg sync.WaitGroup

	// Start worker goroutines
	for range WorkerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(inputPath, chunkChan)
		}()
	}

	// Start reader goroutine
	go func() {
		defer close(chunkChan)
		if err := reader(file, chunkChan); err != nil {
			select {
			case errChan <- err:
			default:
			}
		}
	}()

	// Monitor completion
	go func() {
		wg.Wait()
		close(doneChan)
	}()

	// Wait for completion or error
	select {
	case err := <-errChan:
		return nil, err
	case <-doneChan:
		return &globalContent, nil
	}
}

func reader(file *os.File, chunkChan chan<- Fragment) error {
	buf := make([]byte, BufferSize)
	chunkNum := 0
	currentChunk := make([]byte, 0, ChunkSize)

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		remaining := buf[:n]
		for len(remaining) > 0 {
			available := min(ChunkSize-len(currentChunk), len(remaining))

			currentChunk = append(currentChunk, remaining[:available]...)
			remaining = remaining[available:]

			if len(currentChunk) == ChunkSize || (err == io.EOF && len(currentChunk) > 0) {
				// Make a copy of the chunk data to avoid data races
				chunkData := make([]byte, len(currentChunk))
				copy(chunkData, currentChunk)

				chunkChan <- Fragment{
					Number: chunkNum,
					Data:   chunkData,
				}
				chunkNum++
				currentChunk = currentChunk[:0] // Reset slice
			}
		}
	}
	return nil
}

func worker(basePath string, chunkChan <-chan Fragment) {
	for fragment := range chunkChan {
		chunkPath := fmt.Sprintf("%s.part%d", basePath, fragment.Number)
		fmt.Println("Adding fragment", chunkPath)
		contentMutex.Lock()
		globalContent.Fragments = append(globalContent.Fragments, fragment)
		contentMutex.Unlock()
		/*
			if err := os.WriteFile(chunkPath, chunk.Data, 0644); err != nil {
				select {
				case errChan <- fmt.Errorf("failed to write chunk %d: %w", chunk.Number, err):
				default:
				}
				return
			}
		*/
	}
}
