package main

import "unsafe"

const S = "go" // S[1]-S[0] == 8

func main() {
	var x *[8][8]byte
	println(unsafe.Sizeof((*x)[S[1]-S[0]][S[1]-S[0]]))
}
