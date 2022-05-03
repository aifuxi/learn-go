package main

import "fmt"

// 自动执行，先于main
func init() {
	fmt.Printf("\"init...\": %v\n", "init...")
}

func main() {
	fmt.Printf("\"main...\": %v\n", "main...")
}
