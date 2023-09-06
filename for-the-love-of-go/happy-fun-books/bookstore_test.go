package bookstore_test

import (
	bookstore "happy-fun-books"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuyBook(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}

	want := 1

	result, err := bookstore.Buy(b)

	if err != nil {
		t.Fatal(err)
	}

	got := result.Copies

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestOutOfStock(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}

	_, err := bookstore.Buy(b)

	if err == nil {
		t.Errorf("expected an out of stock error")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	catalog := []bookstore.Book{
		{Title: "For the love of Go"},
		{Title: "The power of Go:Tools"},
	}

	want := catalog
	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
