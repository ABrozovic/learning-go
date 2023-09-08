package main

import (
	"fmt"
)

func main() {
	eggs := 42

	println(eggs)

	eggs = 22

	println(eggs)

	bigger := bigger(2, 1)
	println(bigger)

	greeter := greet("Alice")

	println(greeter)

	println(total([]int{1, 2, 3, 4, 5}))

	evens()

	println(Apply(2, func(x int) int { return x * 2 }))
}

func bigger(param1, param2 int) int {
	if param1 > param2 {
		return param1
	}

	return param2
}

func greet(person string) string {
	switch person {
	case "Alice":
		return fmt.Sprintf("Hello, %s", person)
	case "Bob":
		return fmt.Sprintf("What's up, %s", person)
	}

	return "Hello, stranger"
}

func total(numbers []int) int {
	var total int

	for _, num := range numbers {
		total += num
	}

	return total
}

func evens() {
	var evens []int

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}

	println(fmt.Sprint(evens))
}

func Apply(param int, eval func(param int) int) int {
	return eval(param)
}
