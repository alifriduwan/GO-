package main

import "fmt"

func main() {
	var num1 []int //nil
	if num1 == nil {
		fmt.Println("nil slice")
	}

	num1 = append(num1, 17)
	num1 = append(num1, 17)
	num1 = append(num1, 17)
	fmt.Println(num1)

	num2 := []int{1, 2, 3, 4, 5}
	fmt.Println(num2)

	num3 := num2[1:4] //เป็นการนำข้อมูลใน num2 ตั้งแต่ index1-index3
	fmt.Println(num3)

	num4 := num2[0:] //เป็นการนำข้อมูลใน num2 ตั้งแต่ index0-จนจบ
	fmt.Println(num4)

	num5 := num2[:3]
	fmt.Println(num5) // เป็นการนำข้อมูลใน num2 ตั้งแต่ index0-index2

	num8 := []int{1, 2, 3, 4}
	num9 := num8[1:3] //[2 3]
	fmt.Println(num8) //[1 2 3 4]
	num9[0] = 20      // 2 => 20 และอาเรย์ของ num8 จะเปลี่ยนไปด้วย
	fmt.Println(num8) //[1 20 3 4]
	fmt.Println(num9) //[20 3]

	for i := range num8 { // i will return index of value in num8
		fmt.Println(i)
	}

	for i, v := range num8 { // i will return index of value in num8 and v will retunr value in array num8
		fmt.Println(i, v)
	}

	for _, v := range num8 { // _,v สนใจแค่ค่า v
		fmt.Println(v)
	}

}
