package main

import "fmt"
import "github.com/RanchoCooper/go-cookbook/chapter04/basicerrors"

func main() {

	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
