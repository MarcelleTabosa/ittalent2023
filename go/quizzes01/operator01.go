package main

var x, y = true, false

func o(b bool) bool {
	print(b)
	return !b
}

func main() {
	_ = x || o(x)
	_ = y && o(y)
}
