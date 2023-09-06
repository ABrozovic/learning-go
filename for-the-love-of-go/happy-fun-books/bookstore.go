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

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog []Book, id int) (Book, error) {
	for _, b := range catalog {
		if b.ID == id {
			return b, nil
		}
	}

	return Book{}, errors.New("no book found")
}
