package main

import "fmt"

type Person struct {
	name string
}

// 通过receiver来实现类似类里面的方法
func (per Person) eat() {
	fmt.Printf("name: %v\n", per.name)
}

func main() {
	var p1 Person = Person{
		name: "hello",
	}
	p1.eat()
}
