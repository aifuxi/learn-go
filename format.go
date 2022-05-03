package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	str := "wo shi ni baba"
	fmt.Printf("str: %v\n", str)

	p := Person{Name: "aifuxi", Age: 20}
	fmt.Printf("p: %v\n", p)  // {aifuxi 20} 值
	fmt.Printf("p: %#v\n", p) // main.Person{Name:"aifuxi", Age:20}，相应值的go语法表示
	fmt.Printf("p: %T\n", p)  // 类型

	flag := false
	fmt.Printf("flag: %t\n", flag) // %t 布尔占位
}
