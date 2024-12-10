package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(3, 4)
	point2 := NewPoint(7, 1)

	distance := point1.Distance(point2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
