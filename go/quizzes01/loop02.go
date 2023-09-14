package main

func f() {
	var a = [2]int{5, 7}
	for i, v := range a {
		if i == 0 {
			a[1] = 9
		} else {
			print(v)
		}
	}
}

func g() {
	var a = [2]int{5, 7}
	for i, v := range a[:] {
		if i == 0 {
			a[1] = 9
		} else {
			print(v)
		}
	}
}

func main() {
	f()
	g()
}
