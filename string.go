package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	b := "world"
	fmt.Printf("a + b: %v\n", a+b)

	ab := strings.Join([]string{a, b}, ",")
	fmt.Printf("ab: %v\n", ab)
}
