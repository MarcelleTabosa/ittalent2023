package main

func f(vs ...interface{}) {
	print(len(vs))
}

func main() {
	f()
	f(nil)
	f(nil...)
}
