package main

import (
	"errors"
	"fmt"
)

func Panic() {
	err := errors.New("something error")
	panic(err)
}

func Catcher() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic catched ", r)
		}
	}()
	Panic()
}

func main() {
	fmt.Println("before panic")
	Catcher()
	fmt.Println("after panic")
}
