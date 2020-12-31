package interfaces

import (
	"io"
	"os"
)

// PipeExample helps give some more axamples of using io interfaces
func PipeExample() error {
	// the pipe reader and pipe writer implement io.Reader and io.Writer
	r, w := io.Pipe()

	// this needs to be run in a separate goroutine as it will block waiting for the reader
	// close at the end for cleanup
	go func() {
		_, _ = w.Write([]byte("test\n"))
		_ = w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}

	return nil
}
