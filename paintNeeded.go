package main

import "fmt"

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area / 10.0
}

func main() {
	var amount, total float64
	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("Need: %0.2f l\n", amount)
	total += amount
}
