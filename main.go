package main

import "fmt"

func main() {
	//TH -> Thailand
	//EN -> English
	var langs map[string]string //map[string]string -> [string]->key , string -> [string]=>value
	if langs == nil {
	}
	fmt.Println(langs)

	langs2 := map[string]string{
		"TH": "Thailand",
		"TK": "Turkeys",
		"EN": "English",
	}

	fmt.Println(langs2)

	//การเข้าถึงค่า
	fmt.Println(langs2["TH"])

	//loop
	for k, v := range langs2 {
		fmt.Println(k, ":", v)
	}
}
