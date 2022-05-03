package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 uint8 = 100
	fmt.Printf("num1: %v\n", num1)

	var num2 int8 = -100
	fmt.Printf("num2: %v\n", num2)

	var float1 float32 = 3.14
	fmt.Printf("float1: %v\n", float1)

	var float2 float64 = 3.14
	fmt.Printf("float2: %v\n", float2)

	fmt.Printf("math.min: %v\n", math.MinInt)
}
