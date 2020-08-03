package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer demonstrates some tricks for initializing bytes Buffers
func Buffer(rawString string) *bytes.Buffer {

	rawBytes := []byte(rawString)

	var b = new(bytes.Buffer)
	b.Write(rawBytes)
	// or
	b = bytes.NewBuffer(rawBytes)
	// or
	b = bytes.NewBufferString(rawString)

	return b
}

// ToString is an example of taking an io.Reader and consuming
// it all, then returing a string
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

