package main

import "fmt"

func main() {
	// 定义一个新类型
	type MyInt int
	var num MyInt = 100
	fmt.Printf("num: %T\n", num) // num: main.MyInt

	// 定义一个类型别名
	type MyInt2 = int
	var num2 = 99
	fmt.Printf("num2: %T\n", num2) // num2: int
}
