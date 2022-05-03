package main

import "fmt"

// go里没有继承，可以通过结构体嵌套模拟继承

type Animal struct {
	name  string
	color string
}

type Dog struct {
	Animal
	age int
}

func main() {
	husky := Dog{
		Animal{
			"erha", "black",
		},
		20,
	}

	fmt.Printf("husky: %#v\n", husky)
}
