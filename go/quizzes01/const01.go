package main

const (
	_    = 6
	A, _ = iota, iota + 10
	_, _
	_, B
)

func main() {
	println(A, B)
}
