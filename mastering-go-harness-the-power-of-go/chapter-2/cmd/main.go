package main

import (
	"fmt"
	exercises "mastering-go-exercises/chapter-2"
	"sort"
	"strconv"
	"time"
)

func main() {
	errors()

	r := '€'
	fmt.Println(r)
	fmt.Printf("r %c \n", r)

	loc, err := time.LoadLocation("America/La Paz")

	if err == nil {
		fmt.Println("Full:", loc)
	}

	nums1 := []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	nums2 := []int{11, 12}

	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))

	fmt.Println(nums1, nums2)

	str1 := [5]string{"h", "e", "l", "l", "o"}
	str2 := [5]string{"W", "o", "r", "l", "d"}

	fmt.Print(concat2ArraysToSlice(&str1, &str2))
	fmt.Print(concat2ArraysToArray([3]string{"h", "e", "y"}, [3]string{"y", "e", "a"}))
	fmt.Print(concat2SlicesToSlice([]string{"h", "e", "y"}, []string{"y", "e", "a"}))
}

func concat2ArraysToSlice(arr1, arr2 *[5]string) []string {
	return append(arr1[:], arr2[:]...)
}

func concat2ArraysToArray(arr1, arr2 [3]string) [6]string {
	var newArr [6]string

	copy(newArr[:], arr1[:])
	copy(newArr[len(arr1):], arr2[:])

	return newArr
}

func concat2SlicesToSlice(arr1, arr2 []string) []string {
	return append(arr1, arr2...)
}

func errors() {
	err := exercises.Check(0, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Check() ended normally")
	}

	err = exercises.Check(0, 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Check() ended normally")
	}

	err = exercises.FormattedError(0, 0)

	if err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("-123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Int value is", i)
	}

	i, err = strconv.Atoi("Y123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Int value is", i)
	}
}
