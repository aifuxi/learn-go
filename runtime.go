package main

import (
	"fmt"
	"runtime"
	"time"
)

func showMsg(msg int) {
	fmt.Printf("msg: %v\n", msg)
}

func showMsgPlus() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			// 这句话应该放在被执行的协程里面退出的
			runtime.Goexit()
		}
	}
}

func main() {
	for i := 0; i < 2; i++ {
		go showMsg(i)
	}

	for i := 0; i < 2; i++ {
		// 释放执行资源，让其它协程先执行
		runtime.Gosched()
		fmt.Printf("\"Gosched\": %v\n", "Gosched")
	}

	time.Sleep(time.Second * 2)
	fmt.Printf("\"end...\": %v\n", "end...")
	fmt.Printf("---------\n")
	go showMsgPlus()

	time.Sleep(time.Second * 2)
}
