package main

const X = 3

func main() {
	const (
		X = X + X
		Y
	)

	println(X, Y)
}
