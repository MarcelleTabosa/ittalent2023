package main

type A struct {
	g int
}

func (A) m() int {
	return 1
}

type B int

func (B) g() {}

func (B) f() {}

type C struct {
	A
	B
}

func (C) m() int {
	return 9
}

func main() {
	var c interface{} = C{}
	_, bf := c.(interface{ f() })
	_, bg := c.(interface{ g() })
	i := c.(interface{ m() int })
	println(bf, bg, i.m())
}
