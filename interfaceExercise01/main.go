package main

import "fmt"

type square struct {
	sideLenght float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func main() {
	square := square{
		sideLenght: 10,
	}

	triangle := triangle{
		height: 10,
		base:   10,
	}

	printArea(square)
	printArea(triangle)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
