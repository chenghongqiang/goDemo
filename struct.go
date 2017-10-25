package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books

	Book1.title = "kewin"

	printBook(Book1)

}

func printBook(book Books) {
	fmt.Printf("Book title is : %s\n", book.title)
}

func printBookPoint(book *Books) {

}
