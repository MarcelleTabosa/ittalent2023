package main

func f(n int) (r int) {
	a, r := n-1, n+1
	if a+a == r {
		c, r := n, n*n
		r = r - c
	}
	return r
}

func main() {
	println(f(3))
}
