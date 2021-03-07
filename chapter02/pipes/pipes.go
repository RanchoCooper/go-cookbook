package pipes

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// WordCount takes a file and returns a map with each word as key and it's number of appearances as a value
func WordCount(f io.Reader) map[string]int {
	result := make(map[string]int)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return result
}
