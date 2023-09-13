package exercises

import (
	"crypto/rand"
	"math"
	"math/big"
)

type Shape2d interface {
	Perimeter() float64
}

type Shape3d interface {
	Vol() float64
}

type Circle struct {
	R float64
}

type Cube struct {
	X float64
}

type Cuboid struct {
	X float64
	Y float64
	Z float64
}

type Sphere struct {
	R float64
}

type Shapes3D []Shape3d

func (shapes Shapes3D) Len() int {
	return len(shapes)
}

func (shapes Shapes3D) Less(i, j int) bool {
	return shapes[i].Vol() < shapes[j].Vol()
}

func (shapes Shapes3D) Swap(i, j int) {
	shapes[i], shapes[j] = shapes[j], shapes[i]
}

func (c Cube) Vol() float64 {
	return math.Pow(c.X, 3)
}

func (c Cuboid) Vol() float64 {
	return c.X * c.Y * c.Z
}
func (s Sphere) Vol() float64 {
	return 4 / 3 * math.Pi * math.Pow(s.R, 3)
}

func RandomFloat64(min, max float64) float64 {
	diff := new(big.Float).Sub(new(big.Float).SetFloat64(max), new(big.Float).SetFloat64(min))

	diffInt, _ := diff.Int(nil)
	randInt, err := rand.Int(rand.Reader, diffInt)

	if err != nil {
		return 0
	}

	randFloat := new(big.Float).SetInt(randInt)
	randFloat.Mul(randFloat, diff)
	randFloat.Add(randFloat, new(big.Float).SetFloat64(min))

	randomValue, _ := randFloat.Float64()

	return randomValue
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}
