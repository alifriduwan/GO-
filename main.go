package main

import "fmt"

func Add(n int, m int) int {
	return n + m
}

func Swap(n, m int) (int, int) { //สามารถประกาศแบบนี้เพื่อให้เห็นว่า n and m is integer
	return m, n
}

func main() {
	result1 := Add(17, 14)
	fmt.Println(result1)

	i, j := Swap(1, 2) // 2, 1
	fmt.Println(i, j)

}
