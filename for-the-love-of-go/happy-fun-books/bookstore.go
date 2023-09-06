package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func Buy(book Book) (Book, error) {
	if book.Copies < 1 {
		return book, errors.New("there are no copies left")
	}

	book.Copies--

	return book, nil
}

func GetAllBooks(catalog map[int]Book) map[int]Book {
	return catalog
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	b, ok := catalog[id]
	if !ok {
		return Book{}, errors.New("no book found")
	}

	return b, nil
}
