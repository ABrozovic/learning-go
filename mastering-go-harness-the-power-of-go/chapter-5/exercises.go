package exercises

import "fmt"

func AddFloat(message string, numbers ...float64) float64 {
	fmt.Println(message)

	sum := float64(0)

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func Everything(all ...interface{}) {
	fmt.Println(all...)
}
