package main

const X = '\x61' // 'a'
const Y = 0x62
const A = Y - X // 1
const B int64 = 1

var n = 32

func main() {
	if A == B {
		println(A<<n>>n, B<<n>>n)
	}
}
