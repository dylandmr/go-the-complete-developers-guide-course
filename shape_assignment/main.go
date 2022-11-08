package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	sq := square{sideLength: 12}
	tr := triangle{base: 4, height: 7}

	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
