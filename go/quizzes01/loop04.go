package main

func main() {
	i, s := 9, []int{}

	for i = range s {
	}
	print(i)

	for i = 0; i < len(s); i++ {
	}
	print(i)

	s = append(s, 1, 2, 3, 4, 5)

	for i = range s {
	}
	print(i)

	for i = 0; i < len(s); i++ {
	}
	println(i)
}
