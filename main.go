package main

import (
	"fmt"
)

func main() {
	fmt.Println(1)
	defer fmt.Println(2) //ทำก็ต่อเมื่อฟังก์ชันทำงานเสร็จแล้ว
	fmt.Println(3)
}
