package main

const N = 1

var n = N
var a byte = 128 << N >> N
var b byte = 128 << n >> n

func main() {
	println(a, b)
}
