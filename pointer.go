package main

import "fmt"

func main() {
	/*
		&取地址值，取到的地址值可以复制给指针变量
		*是用来取值的
	*/
	var p1 *int
	var num = 100
	p1 = &num
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p1 value: %v\n", *p1)

	arr := [3]int{1, 2, 3}
	var arrPointer [3]*int

	fmt.Printf("arrPointer: %v\n", arrPointer)

	for i := 0; i < len(arr); i++ {
		arrPointer[i] = &arr[i]
	}

	fmt.Printf("arrPointer: %v\n", arrPointer)

	for _, v := range arrPointer {
		fmt.Printf("v: %v\n", *v)
	}
}
