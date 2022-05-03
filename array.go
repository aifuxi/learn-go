package main

import "fmt"

func main() {
	var numArr = [3]int{10, 20, 30}
	fmt.Printf("numArr: %v\n", numArr)

	// ...不指定长度，自动推断
	var numArr2 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("numArr2: %v\n", numArr2)
	fmt.Printf("%v\n", len(numArr2))
}
