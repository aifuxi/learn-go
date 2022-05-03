package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// go xxx开启一个协程
	go showMsg("go routine")
	showMsg("doooo doo")

	time.Sleep(time.Millisecond * 1000)
}
