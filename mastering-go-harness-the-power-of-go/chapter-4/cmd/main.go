package main

import (
	"fmt"
	exercises "mastering-go-exercises/chapter-4"
	"sort"
)

func main() {
	CustomSort()
}

func CustomSort() {
	dataSet := make(exercises.MyStructs, 5)

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			dataSet[i] = exercises.ShouldOrder{Order: i}
		} else {
			dataSet[i] = exercises.ShouldDifferentiate{Diff: false}
		}
	}

	for i, data := range dataSet {
		switch data.(type) {
		case exercises.ShouldDifferentiate:
			fmt.Println(i, "should differentiate")
		case exercises.ShouldOrder:
			fmt.Println(i, "should Order")
		}
	}

	sort.Sort(sort.Reverse(dataSet))
	fmt.Println(dataSet)
}

func Shape3D() {
	data := exercises.Shapes3D{}

	for i := 0; i < 3; i++ {
		cube := exercises.Cube{X: exercises.RandomFloat64(1, 3)}
		cuboid := exercises.Cuboid{X: exercises.RandomFloat64(1, 3), Y: exercises.RandomFloat64(1, 3), Z: exercises.RandomFloat64(1,
			3)}
		sphere := exercises.Sphere{R: exercises.RandomFloat64(1, 3)}

		data = append(data, cube, cuboid, sphere)
	}

	PrintShapes(data)
	sort.Sort(data)
	PrintShapes(data)
	sort.Sort(sort.Reverse(data))
	PrintShapes(data)
}

func PrintShapes(a exercises.Shapes3D) {
	for _, v := range a {
		switch v.(type) {
		case exercises.Cube:
			fmt.Printf("Cube: volume %.2f\n", v.Vol())
		case exercises.Cuboid:
			fmt.Printf("Cuboid: volume %.2f\n", v.Vol())
		case exercises.Sphere:
			fmt.Printf("Sphere: volume %.2f\n", v.Vol())
		default:
			fmt.Println("Unknown data type!")
		}
	}

	fmt.Println()
}

func Shape2D() {
	circle := exercises.Circle{R: 1.5}

	fmt.Printf("R %.2f -> Perimeter %.3f \n", circle.R, circle.Perimeter())

	_, ok := interface{}(circle).(exercises.Shape2d)

	if ok {
		fmt.Println("Circle is a Shape2d")
	}

	fmt.Println(ok)
}
