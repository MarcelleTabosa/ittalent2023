package main

func main() {
	x := []int{7, 8, 9}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			print(v)
		}()
		y[i] = &i
	}
	print(*y[0], *y[1], *y[2], " ")
}
