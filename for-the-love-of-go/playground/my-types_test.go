package playground_test

import (
	"playground"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()

	want := playground.MyInt(18)
	got := playground.MyInt(9).Twice()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()

	input := playground.MyString("hello")

	want := 5
	got := input.MyStringLen()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
