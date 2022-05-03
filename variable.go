package main

import "fmt"

// 匿名变量
func getInfo() (string, int) {
	return "aifuix", 10
}

func main() {
	var str string = "aifuxi"
	var str2 = "This is a str2"
	fmt.Printf("str: %v\n", str)
	fmt.Printf("str2: %v\n", str2)

	var flag bool = true
	fmt.Printf("flag: %v\n", flag)

	// 短变量类型声明，只能用在函数内部
	num := 100
	fmt.Printf("num: %v\n", num)

	// 匿名变量
	_, age := getInfo()
	fmt.Printf("age: %v\n", age)
}
