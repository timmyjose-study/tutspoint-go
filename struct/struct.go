package main

import "fmt"

type book struct {
	title  string
	author string
	id     int
}

func processBook(b *book) {
	fmt.Printf("Title: %s, author: %s, id: %d\n", b.title, b.author, b.id)
}

func main() {
	b1 := book{title: "War and Peace", author: "Leo Tolstoy", id: 1}
	processBook(&b1)

}