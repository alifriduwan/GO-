package main

import "fmt"

func main() {
	i := 10
	j := &i
	i = 12
	*j = 15
	fmt.Println(*j, i)
}
