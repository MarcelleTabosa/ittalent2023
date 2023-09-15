package main

import (
	"fmt"
)

type Person struct {
	name     string
	lastName string
}

type No struct {
	person   Person
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

func (l *List) addData(data Person) {
	var no No
	no.person.name = data.name
	no.person.lastName = data.lastName

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
			"Current position: %p %d - data: %s - previous: %p - next %p - deleted: %t\n",
			no, no.position, no.person, no.previous, no.next, no.inactive)
		no = no.next
	}
}

func (l List) findByData(data Person) []uint {
	var result []uint
	var no *No

	no = l.start
	for no != nil {
		if no.person == data {
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
	no = l.start
	for no != nil {
		if no.inactive {
			deleted = no
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

			for deleted != nil {
				deleted.position = deleted.position - 1
				deleted = deleted.next
			}
			deleted = nil
		}
		no = no.next
	}
}

func (l *List) updatePerson(position uint, data Person) {
	var no *No
	no = l.start

	for no != nil {
		if no.position == position {
			no.person.name = data.name
			no.person.lastName = data.lastName
			no = nil
		} else {
			no = no.next
		}
	}
}

func main() {
	var list List
	list.addData(Person{"Marcelle", "Oliveira"})
	list.addData(Person{"Marcelle", "Tabosa"})
	list.addData(Person{"Marcelle", "Souza"})
	list.addData(Person{"Marcelle", "Souza"})
	list.addData(Person{"Marcelle", "Oliveira"})
	list.addData(Person{"Marcelle", "Tabosa"})
	list.addData(Person{"Marcelle", "Oliveira"})
	list.showList()
	positions := list.findByData(Person{"Marcelle", "Oliveira"})
	fmt.Println(positions)
	nos := list.findByPosition(positions)
	fmt.Println(nos)
	deletedes := list.logicalDeletion(positions)
	fmt.Println(deletedes)
	list.showList()
	fmt.Println("Deletion defined")
	list.deletionDefined()
	list.showList()
	list.updatePerson(0, Person{"Marcelle", "Tabosa de Souza"})
	list.showList()
}
