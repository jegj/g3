package chunkio

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func Splitter(filePath string, chunkSize int, numWorkers int, processor ChunkProcessor) error {
	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stats: %w", err)
	}
	fileSize := fileInfo.Size()

	type chunkWork struct {
		offset     int64
		size       int
		chunkIndex int
	}
	workChan := make(chan chunkWork, numWorkers)
	var wg sync.WaitGroup
	errChan := make(chan error, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			buffer := make([]byte, chunkSize)

			for work := range workChan {
				sectionReader := io.NewSectionReader(file, work.offset, int64(work.size))
				
				n, err := sectionReader.Read(buffer[:work.size])
				if err != nil && err != io.EOF {
					errChan <- fmt.Errorf("error reading chunk %d: %w", work.chunkIndex, err)
					return
				}
				
				if n > 0 {
					if err := processor(buffer[:n], work.chunkIndex); err != nil {
						errChan <- fmt.Errorf("error processing chunk %d: %w", work.chunkIndex, err)
						return
					}
				}
			}
		}()
	}

	// Queue work items
	chunkIndex := 0
	for offset := int64(0); offset < fileSize; offset += int64(chunkSize) {
		size := chunkSize
		if offset+int64(chunkSize) > fileSize {
			size = int(fileSize - offset)
		}
		
		workChan <- chunkWork{
			offset:     offset,
			size:       size,
			chunkIndex: chunkIndex,
		}
		chunkIndex++
	}
	close(workChan)

	wg.Wait()
	close(errChan)

	select {
	case err := <-errChan:
		return err
	default:
		return nil
	}
}

