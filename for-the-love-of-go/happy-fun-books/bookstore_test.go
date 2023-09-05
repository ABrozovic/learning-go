package bookstore_test

import (
	bookstore "happy-fun-books"
	"testing"
)

func testBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title: "Spark Joy",
		Author: "Marie  Kondo",
		Copies: 2,
	}

}
