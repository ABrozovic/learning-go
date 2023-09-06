package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNewCreditCard(t *testing.T) {
	t.Parallel()

	want := "123456"

	c, err := creditcard.New(want)
	got := c.Number()

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestInvalidCreditcardNumber(t *testing.T) {
	_, err := creditcard.New("")

	if err == nil {
		t.Fatalf("want error for empty number")
	}
}
