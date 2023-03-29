package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrorInvalidRadius = errors.New("radius must be positive")

func Area(r float64) (float64, error) {
	if r <= 0 {
		return 0.0, ErrorInvalidRadius
	}

	return math.Pi * r * r, nil
}

func main() {
	result, err := Area(-17)
	if errors.Is(err, ErrorInvalidRadius) {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
