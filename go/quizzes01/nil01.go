package main

var x = *new(*int)
var y *int = nil

func f() interface{} {
	return y
}

func main() {
	if f() == nil {
		if x == nil {
			println("A")
		} else {
			println("B")
		}
	} else {
		if x == nil {
			println("C")
		} else {
			println("D")
		}
	}
}
