package main

import "fmt"

type No struct {
	data     int
	previous *No
	next     *No
	position uint
}

type List struct {
	start *No
	end   *No
}

func (l List) setLastPosition() uint {
	position := l.end.position + 1
	return position
}

func (l *List) addData(data int) {
	var no No
	no.data = data
	if l.start == nil {
		l.start = &no
		l.end = &no
	} else {
		no.position = l.setLastPosition()
		no.previous = l.end
		l.end.next = &no
		l.end = &no
	}
}

func (l List) showList() {
	var no *No
	no = l.start
	fmt.Println("List:")
	for no != nil {
		fmt.Printf(
			"Current position: %p - data: %d - previous: %p - next %p\n",
			no, no.data, no.previous, no.next)
		no = no.next
	}
}

func main() {
	var list List
	list.addData(45)
	list.addData(55)
	list.addData(88)
	list.addData(99)

	list.showList()
}
