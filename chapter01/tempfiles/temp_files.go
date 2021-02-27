package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

// WorkWithTemp will give some basic patterns for working with temporary files and directories
func WorkWithTemp() error {
	// the first argument being blank means it will use create the directory in the location returned by os.TempDir()
	fmt.Println(os.TempDir())
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}
	defer os.RemoveAll(t)

	// the directory must exists to create the tempfile created
	// t is an *os.File object
	tf, err := ioutil.TempFile(t, "tmp.txt")
	if err != nil {
		return err
	}
	fmt.Println(tf.Name())

	return nil
}
