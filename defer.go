package main

import "fmt"

func main() {
	// output
	/*
	   "start": start
	   "end": end
	   step 3
	   step 2
	   step 1
	*/
	fmt.Printf("\"start\": %v\n", "start")
	// defer， 延迟执行，等到函数return时再倒叙执行defer后的语句，
	defer fmt.Printf("step 1\n")
	defer fmt.Printf("step 2\n")
	defer fmt.Printf("step 3\n")
	fmt.Printf("\"end\": %v\n", "end")

}
