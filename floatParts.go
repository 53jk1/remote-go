package main

import (
	"fmt"
	"math"
)

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	var inputNumber float64
	fmt.Scanln(&inputNumber)
	cans, remainder := floatParts(inputNumber)
	fmt.Println(cans, remainder)
}
