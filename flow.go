package main

import "fmt"

func main() {
	score := 100

	if score < 60 {
		println("不及格")
	} else if score >= 60 && score <= 90 {
		println("良好")
	} else {
		println("优秀")
	}

	arr := [...]int{100, 2321, 11, 45, 67, 22}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("i: %v\n", arr[i])
	}

	fmt.Printf("---\n")

	for _, v := range arr {
		fmt.Printf("v: %v\n", v)
	}
}
