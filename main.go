package main

import (
	"fmt"
)

func main() {
	// n := 17
	if n := 17; n%2 == 0 { //สารมารถกำหนดค่าเริ่มต้นตรงตำแหน่งเช็คเงื่อนไข
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}
