package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer fmt.Print(recover())
		defer panic(1)
		recover()
	}()
	defer recover()
	panic(2)
}
