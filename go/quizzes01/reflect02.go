package main

import "reflect"

type T struct {
	A int
	b int
}

func (T) M() {}
func (T) m() {}

func main() {
	v := reflect.ValueOf(T{})
	println(v.NumField(), v.NumMethod())
}
