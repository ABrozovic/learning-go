package playground

type MyString string

type MyInt int

func (s MyString) MyStringLen() int {
	return len(s)
}

func (i MyInt) Twice() MyInt {
	return i * 2
}
