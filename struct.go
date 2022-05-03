package main

import "fmt"

func main() {

	type Person struct {
		id   string
		name string
		age  int
	}

	var p1 Person
	p1.name = "aifuxi"
	p1.age = 24
	fmt.Printf("p1: %#v\n", p1)

	// 匿名结构体
	var dog struct {
		color string
	}
	dog.color = "red"
	fmt.Printf("dog: %#v\n", dog)

	// 结构体初始化
	p2 := Person{
		id:   "xxx",
		name: "aifuxi",
		age:  20,
	}

	fmt.Printf("p2: %v\n", p2)
}
