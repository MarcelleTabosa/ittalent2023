package main

func bar() (r int) {
	defer func() {
		r += 4
		if recover() != nil {
			r += 8
		}
	}()

	var f func()
	defer f()
	f = func() {
		r += 2
	}

	return 1
}

func main() {
	println(bar())
}
