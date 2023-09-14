package main

type Foo struct {
	v int
}

func MakeFoo(n *int) Foo {
	print(*n)
	return Foo{}
}

func (Foo) Bar(n *int) {
	print(*n)
}

func main() {
	var x = 1
	var p = &x
	defer MakeFoo(p).Bar(p)
	x = 2
	p = new(int)
	MakeFoo(p)
}
