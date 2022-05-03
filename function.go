package main

import "fmt"

func sum(a int, b int) (c int) {
	c = a + b
	return c
}

func foo(a int, b int) int {
	return a + b
}

func bar(a int, b int) (c int, d int) {
	c = 100
	d = 200

	return // 相当于 return c, d
}

// 定义一个f1的函数类型
type f1 func(int, int) int

// 可变参数
func fn(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	a := sum(1, 2)
	fmt.Printf("a: %v\n", a)

	fn(1, 2, 3, 4, 6)

	var fn1 f1
	fn1 = sum
	b := fn1(100, 200)
	fmt.Printf("b: %v\n", b)
}
