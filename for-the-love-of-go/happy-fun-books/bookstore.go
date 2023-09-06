package bookstore

import "errors"

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

type Catalog map[int]Book

func Buy(book Book) (Book, error) {
	if book.Copies < 1 {
		return book, errors.New("there are no copies left")
	}

	book.Copies--

	return book, nil
}

func (c Catalog) GetAllBooks() []Book {
	books := []Book{}
	for _, b := range c {
		books = append(books, b)
	}

	return books
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	b, ok := catalog[id]
	if !ok {
		return Book{}, errors.New("no book found")
	}

	return b, nil
}

func PriceCents(b Book) int {
	return b.PriceCents - (b.PriceCents * b.DiscountPercent / 100)
}

func (b Book) NetPriceCents() int {
	totalDiscount := (b.PriceCents * b.DiscountPercent / 100)
	return b.PriceCents - totalDiscount
}
