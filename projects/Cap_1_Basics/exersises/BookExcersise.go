package main

import (
	"Basics/structs"
	"fmt"
)

func AddBook(book structs.Book, books []structs.Book) []structs.Book {

	fmt.Print("Adding book to slice")

	books = append(books, book)
	return books

}

func RemoveBook(book structs.Book, books []structs.Book) []structs.Book {

	fmt.Print("Removing book to slice")

	return books
}
