package main

import (
	exercises "mastering-go-exercises/chapter-1"
	"os"
)

func main() {
	arguments := os.Args
	exercises.PhoneBook(arguments...)
}

// exercises.SysLog()
// exercises.CustomLog()
// exercises.Which(arguments...)
