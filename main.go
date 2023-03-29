package main

import (
	"errors"
	"fmt"
	"math"
)

func Area(r float64) (float64, error) {
	if r <= 0 {
		return 0.0, errors.New("radius must be positive")
	}

	return math.Pi * r * r, nil
}

func main() {
	result, err := Area(-17)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
