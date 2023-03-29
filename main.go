package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) AreaCircle() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Point struct {
	X int
	Y int
}

func main() {
	c := Circle{17}
	a := c.AreaCircle()
	fmt.Println(c, a)
	p := Point{X: 14, Y: 17}
	fmt.Println(p)
}
