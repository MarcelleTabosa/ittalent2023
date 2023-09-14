package main

func main() {
	count := 0
	for i := range [256]struct{}{} {
		if n := byte(i); n == -n {
			count++
		}
	}
	println(count)
}
