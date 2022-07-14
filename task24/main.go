package main

import (
	"fmt"
	"math"
)

// Point represent coordinate point
type Point struct {
	x, y float64
}

// X is a getter for x
func (p Point) X() float64 {
	return p.x
}

// Y is a getter for y
func (p Point) Y() float64 {
	return p.y
}

//New is a constructor func for Point type
func New(x, y float64) *Point {
	return &Point{x, y}
}

// Distance method returns distance between two points
func (p Point) Distance(that *Point) float64 {
	// it implements mathematical formula
	x, y := p.x-that.x, p.y-that.y
	return math.Sqrt(x*x + y*y)
}

func main() {
	p1 := New(3.0, 12.0)
	p2 := New(-11.0, 1.0)

	fmt.Println(p1.Distance(p2))
}
