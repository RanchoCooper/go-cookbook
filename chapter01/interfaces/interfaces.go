package interfaces

import (
	"fmt"
	"io"
	"os"
)

func Copy(in io.ReadSeeker, out io.Writer) error {

	// write result to `out` and `Stdout`
	w := io.MultiWriter(out, os.Stdout)

	// it's dangerous if the incoming data is too large
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	// PENDING why?
	in.Seek(0, 0)
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	fmt.Println()

	return nil
}