package playground

import (
	"errors"
	"strings"
)

type MyString string

type MyInt int

type MyStringBuilder strings.Builder

type MyBuilder struct {
	Contents strings.Builder
}

func (s MyString) Len() int {
	return len(s)
}

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (mySb MyStringBuilder) Hello() string {
	return "Hello"
}

func (s MyString) ToUpper() string {
	return strings.ToUpper(string(s))
}

func (i *MyInt) Double() {
	*i *= 2
}

func Withdraw(amount int) (int, error) {
	balance := 1000

	if balance >= amount {
		return balance - amount, nil
	}

	return 0, errors.New("not enough funds")
}
