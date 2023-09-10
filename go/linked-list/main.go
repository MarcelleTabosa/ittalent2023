package main

import (
	"fmt"
)

type No struct {
	data     int
	position uint
	inactive bool
	previous *No
	next     *No
}

type Deleted struct {
	position uint
	deleted  bool
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
			"Current position: %p %d - data: %d - previous: %p - next %p - deleted: %t\n",
			no, no.position, no.data, no.previous, no.next, no.inactive)
		no = no.next
	}
}

func (l List) findByData(data int) []uint {
	var result []uint
	var no *No

	no = l.start
	for no != nil {
		if no.data == data {
			result = append(result, no.position)
		}
		no = no.next
	}
	return result
}

func (l List) findByPosition(positions []uint) []No {

	var result []No
	var no *No

	for _, position := range positions {
		no = l.start
		for no != nil {
			if no.position == position {
				result = append(result, *no)
			}
			no = no.next
		}
	}
	return result
}

func (l *List) logicalDeletion(positions []uint) []Deleted {
	var deleted []Deleted
	var no *No

	no = l.start
	for no != nil {
		for _, position := range positions {
			if no.position == position {
				no.inactive = true

				var n Deleted
				n.position = position
				n.deleted = no.inactive

				deleted = append(deleted, n)
			}
		}
		no = no.next
	}
	return deleted
}

func (l *List) deletionDefined() {
	var no, deleted *No
	var position uint = 0

	no = l.start
	for no != nil {
		if no.inactive {
			if no == l.start {
				l.start = no.next
				l.start.previous = nil
			} else if no.next == nil {
				l.end = no.previous
				l.end.next = nil
			} else {
				no.previous.next = no.next
				no.next.previous = no.previous
			}
			deleted = no
			position = no.position
			for deleted != nil {
				socorroDeus := deleted.next.position
				deleted.next.position = position
				position = socorroDeus
				deleted = deleted.next
			}
		}
		no = no.next
	}
}

func main() {
	var list List
	list.addData(10)
	list.addData(55)
	list.addData(26)
	list.addData(88)
	list.addData(10)
	list.addData(34)
	list.addData(10)
	list.showList()
	positions := list.findByData(10)
	fmt.Println(positions)
	nos := list.findByPosition(positions)
	fmt.Println(nos)
	deletedes := list.logicalDeletion(positions)
	fmt.Println(deletedes)
	list.showList()
	fmt.Println("Deletion defined")
	list.deletionDefined()
	list.showList()
}
