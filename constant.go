package main

import "fmt"

func main() {
	const PI float32 = 3.14
	fmt.Printf("PI: %v\n", PI)

	const MALE = 1
	fmt.Printf("MALE: %v\n", MALE)

	const name, age, gender = "aifuix", 24, MALE
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("gender: %v\n", gender)

	const (
		width  = 100
		height = 200
	)
	fmt.Printf("width: %v\n", width)
	fmt.Printf("height: %v\n", height)

	// iota， go内置的一个可变常量，每声明一个常量 iota +1从0开始
	const (
		a1 = iota // 0
		a2 = iota // 1
		a3 = iota // 2
		_         // 跳过，3
		a4 = iota // 4
		a5 = 100  // 跳过，100
		a6 = iota // 101
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)
	fmt.Printf("a5: %v\n", a5)
	fmt.Printf("a6: %v\n", a6)
}
