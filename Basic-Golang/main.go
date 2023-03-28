package main //จุดเริ่มต้นการทำงาน
import "fmt"

// ชนิดข้อมูล

//string ถ้าตัวแปรไม่มีการกำหนดตัวแปร ค่าเริ่มต้นจะเท่ากับ ''

//bool	//ถ้าตัวแปรไม่มีการกำหนดตัวแปร ค่าเริ่มต้นจะเท่ากับ false

//int int8 int16 int32 int64
//uint uint8 uint16 uint32 uintptr

// float32 float64

//เมื่อโปรแกรมรัน จะมาทำงานที่ฟังก์ชัน main เป็นลำดับแรก
func main() {
	var num int //ถ้าตัวแปรไม่มีการกำหนดตัวแปร ค่าเริ่มต้นจะเท่ากับ 0
	var sum int = 17
	var flag bool
	// var age int = 15
	age := 15
	i, j, k := 1, 2, 3
	pi := 3.14
	// convert float to uint
	fmt.Println(pi, uint(pi))
	fmt.Println("Hello world")
	fmt.Println(num, sum, i, j, k, age, flag)
}
