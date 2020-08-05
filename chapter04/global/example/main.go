package main

import "github.com/RanchoCooper/go-cookbook/chapter04/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
