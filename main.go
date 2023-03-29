package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	//process file
	//read file
	//calculate...
}
