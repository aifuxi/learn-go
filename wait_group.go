package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func showMsg(msg int) {
	// 函数结束时，一个标记done
	defer wp.Done()
	fmt.Printf("msg: %v\n", msg)
}

func main() {

	for i := 0; i < 10; i++ {
		go showMsg(i)
		// 启动协程时，WaitGroup添加标记
		wp.Add(1)
	}

	// 等待所有标记结束
	wp.Wait()
	fmt.Printf("\"end...\": %v\n", "end...")
}
