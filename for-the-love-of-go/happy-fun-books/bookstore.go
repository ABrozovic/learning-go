package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(book Book) (Book, error) {
	if book.Copies < 1 {
		return book, errors.New("there are no copies left")
	}

	book.Copies--

	return book, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}
