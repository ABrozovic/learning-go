package playground_test

import (
	"playground"
	"strings"
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
	got := input.Len()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestStringBuilder(t *testing.T) {
	t.Parallel()

	var sb strings.Builder

	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")

	want := "Hello, Gophers!"
	got := sb.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

	wantlen := 15

	gotLen := sb.Len()

	if wantlen != gotLen {
		t.Errorf("%q: want len %d, got len %d", sb.String(), wantlen, gotLen)
	}
}

func TestMyStringBuilderHello(t *testing.T) {
	t.Parallel()

	want := "Hello"

	got := playground.MyStringBuilder{}.Hello()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestMyBuilder(t *testing.T) {
	t.Parallel()

	var mb playground.MyBuilder

	mb.Contents.WriteString("Hello, ")
	mb.Contents.WriteString("Gophers!")

	want := "Hello, Gophers!"
	got := mb.Contents.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStringUpperCase(t *testing.T) {
	t.Parallel()

	want := "HELLO"
	got := playground.MyString("hello").ToUpper()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()

	got := playground.MyInt(12)
	want := playground.MyInt(24)

	p := &got
	p.Double()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
