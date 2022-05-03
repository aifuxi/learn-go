package main

import "fmt"

func main() {
	var m1 map[string]string
	m1 = make(map[string]string)
	m1["name"] = "aifuxi"
	m1["age"] = "age"
	fmt.Printf("m1: %v\n", m1)

	value, ok := m1["name"] // ok是bool,表示value是否存在
	fmt.Printf("value: %v\n", value)
	fmt.Printf("ok: %v\n", ok)

	for key, value := range m1 {
		fmt.Printf("%v: %v \n", key, value)
	}
}
