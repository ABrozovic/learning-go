package main

import (
	"calculator"
	"fmt"
)

func main() {
	var testData float64 = 2
	result := calculator.Add(testData, testData)
	fmt.Printf("This is the result %f", result)
}
