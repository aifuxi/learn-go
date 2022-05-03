package main

import "fmt"

func main() {
	str := "abc"
	fmt.Printf("%T\n", str)

	flag := false
	fmt.Printf("%T\n", flag)

	num := 100
	fmt.Printf("num: %T\n", num)

	// 数组
	arr := [5]int{1, 2, 3, 99, 20}
	fmt.Printf("arr: %T\n", arr)

	// 切片，可以看成动态的数组
	slice := []string{"aaa", "bbbb"}
	fmt.Printf("slice: %T\n", slice)
}
