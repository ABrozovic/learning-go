package main

import (
	"fmt"
	exercises "mastering-go-exercises/chapter-1"
	"os"
)

func main() {
	arguments := os.Args
	err := exercises.Which(arguments...)

	if err != nil {
		fmt.Println(err)
	}
}

// arguments := os.Args
// exercises.PhoneBook(arguments...)
// exercises.SysLog()
// exercises.CustomLog()
// exercises.Which(arguments...)
