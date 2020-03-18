package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

func getType(s shape) string {
	switch s.(type) {
	case square:
		return "square"
	case circle:
		return "circle"
	default:
		return "unknown geometrical shape"
	}
}

func info(s shape) {
	fmt.Println("Area of ", s.area(), "for geometrical shape", getType(s))
}

type shape interface {
	area() float64
}

func main() {
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}

	info(circ)
	info(squa)

	func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}()
}
