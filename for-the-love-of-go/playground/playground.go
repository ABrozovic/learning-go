package playground

import "strings"

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
