package main

import (
	"course-go/shape" //import module and package
	"fmt"
)

func main() {
	area := shape.GetCircleArea(10)
	fmt.Print("Circle Area:", area)
}
