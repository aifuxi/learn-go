package main

import (
	"fmt"
	"strings"
)

func add1(r rune) rune {
	return r + 1
}

func main() {
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "Admix"))

	// 使用匿名函数
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000"))
}
