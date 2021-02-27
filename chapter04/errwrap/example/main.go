package main

import (
	"fmt"

	"github.com/RanchoCooper/go-cookbook/chapter04/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}
