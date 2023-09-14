package main

type Book struct {
	Pages int
}

func f() int {
	var books = []Book{{555}}
	for _, book := range books {
		book.Pages = 999
	}
	return books[0].Pages
}

func g() int {
	var books = []*Book{{555}}
	for _, book := range books {
		book.Pages = 999
	}
	return books[0].Pages
}

func main() {
	println(f(), g())
}
