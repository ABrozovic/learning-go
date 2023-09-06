package bookstore

import "errors"

type Category int

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
	CategorySciFi
)

var isValidCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
	CategorySciFi:             true,
}

type Book struct {
	Title           string
	Author          string
	category        Category
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

type Catalog map[int]Book

func Buy(book *Book) (*Book, error) {
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

func PriceCents(b *Book) int {
	return b.PriceCents - (b.PriceCents * b.DiscountPercent / 100)
}

func (b *Book) NetPriceCents() int {
	totalDiscount := (b.PriceCents * b.DiscountPercent / 100)
	return b.PriceCents - totalDiscount
}

func (b *Book) SetPriceCents(price int) error {
	if price < 1 {
		return errors.New("the price can't be 0 or lower")
	}

	b.PriceCents = price

	return nil
}

func (b *Book) SetCategory(cat Category) error {
	if !isValidCategory[cat] {
		return errors.New("there is no such category")
	}

	b.category = cat

	return nil
}

func (b *Book) Category() Category {
	return b.category
}
