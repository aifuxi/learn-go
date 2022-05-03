package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		Name: "aifuxi",
		Age:  24,
	}
	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))

	b2 := []byte(`{"Name":"aifuxi","Age":24}`)
	var f interface{}
	json.Unmarshal(b2, &f)
	fmt.Printf("f: %v\n", f)
}
