package main

import (
	"flag"
	"fmt"
	"github.com/RanchoCooper/go-cookbook/chapter02/flags"
)

func main() {
	c := flags.Config{}
	c.Setup()

	// generally call this from main
	flag.Parse()

	fmt.Println(c.GetMessage())
}
