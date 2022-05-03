package main

import "fmt"

func main() {
	// 可以在if语句里定义变量，该变量作用域是这个if语句里，超出不能使用
	// 判断条件里只能是bool或者bool类型的表达式，不能是非0
	if a := 16; a > 18 {
		fmt.Printf("已成年  %v\n", a)
	} else {
		fmt.Printf("未成年  %v\n", a)
	}

	// b := 20

	/* if b {

	} */

}
