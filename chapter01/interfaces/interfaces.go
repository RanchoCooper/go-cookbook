package interfaces

import (
	"fmt"
	"io"
	"os"
)

// Copy copies data from in to out first directly,
// then using a buffer to write. It also writes to stdout
func Copy(in io.ReadSeeker, out io.Writer) error {

	// write to out, but also Stdout
	w := io.MultiWriter(out, os.Stdout)

	// a standard copy
	// NOTE: this can be dangerous if there's a lot of data in in
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	in.Seek(0, 0)

	// buffered write
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	fmt.Println()
	return nil
}