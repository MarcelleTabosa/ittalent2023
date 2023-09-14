package main

import (
	"fmt"
	"reflect"
)

var x = []int{2: 5, 6, 0: 7}

func main() {
	fmt.Println(x)
	for i, item := range x {
		fmt.Println(i, item)
	}
	fmt.Println(reflect.TypeOf(x[1]))
	fmt.Println(reflect.TypeOf(x))
}
