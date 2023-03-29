package main

import "fmt"

func main() {
	// var num1 [2]int
	// fmt.Println(num1) // [0 0] เพราะไม่มีการกำหนดค่า

	var num1 [2]int
	num1[0] = 1
	num1[1] = 2
	fmt.Println(len(num1), num1, num1[0], num1[1])

	num2 := [3]int{14, 17, 15}
	fmt.Println(len(num2), num2)
}
