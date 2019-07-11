package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLenght float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	quadrato := square{2}
	triangolo := triangle{2, 1}

	printArea(quadrato)
	printArea(triangolo)
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) * 0.5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
