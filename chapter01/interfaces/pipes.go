package interfaces

import (
	"io"
	"os"
)

// PipeExample helps give some more axamples of using interfaces
func PipeExample() error {

	r, w := io.Pipe()

	// this needs to be run in a separate goroutine as it will block waiting for the reade.
	// remember to close at the end of cleanup
	go func() {
		w.Write([]byte("test\n"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}

	return nil
}
