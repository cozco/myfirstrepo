package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	mySquare := square{
		sideLength: 10.0,
	}
	myTriangle := triangle{
		height: 15.0,
		base:   25.0,
	}

	printArea(mySquare)
	printArea(myTriangle)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}
