package main

import (
	"fmt"
	"strings"
)

type Printer interface {
	Print() string //method Print() return string
}
type PDF struct {
	Content string
}

func (p PDF) Print() string {
	return p.Content
}

type Excel struct {
	Cells []string
}

func (e Excel) Print() string {
	return strings.Join(e.Cells, ",")
}

func main() {
	docs := []Printer{
		PDF{Content: "Hello PDF"},
		Excel{Cells: []string{"A", "B"}},
	}

	for _, doc := range docs {
		fmt.Println(doc.Print())
	}
}
