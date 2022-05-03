package main

import "fmt"

func main() {
	var slice1 []int
	var slice3 = []int{10, 20, 30}
	fmt.Printf("slice1: %v\n", slice1)
	// nil 空，切片初始化时，容量为0长度为0
	fmt.Printf("slice1: %t\n", slice1 == nil)
	// len
	fmt.Printf("len: %v\n", len(slice1))
	slice1 = slice3[1:3]

	// make创建slice时，已经先分配好了一个内存了
	var slice2 = make([]int, 5)
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("len: %v\n", len(slice1))
	fmt.Printf("le2: %v\n", cap(slice1))
	fmt.Printf("cap: %v\n", cap(slice2))
	slice2 = slice3[1:3]
	fmt.Printf("cap2: %v\n", cap(slice2))

	fmt.Printf("-------------\n")
	// append
	var slice4 []int
	var slice5 = make([]int, 4)
	slice4 = append(slice4, 111, 222, 333)
	fmt.Printf("slice4: %v\n", slice4) //slice4: [111 222 333]
	// delete  删除index元素，append(slice4[:index], slice4[index+1:]...)
	slice4 = append(slice4[:1], slice4[2:]...)
	fmt.Printf("slice4: %v\n", slice4)
	// copy复制时，需要使用make创建的slice
	copy(slice5, slice4)
	// update
	slice4[1] = 99999999
	fmt.Printf("slice4: %v\n", slice4)

}
