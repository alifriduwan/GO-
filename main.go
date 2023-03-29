package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	i := 10
	j := &i
	i = 12
	*j = 15
	fmt.Println(*j, i)

	p := &Point{X: 100, Y: 200}
	(*p).X = 300
	fmt.Println(p)
}
