package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var a int32 = 100

func add() {
	// cas， newValue oldValue
	atomic.AddInt32(&a, 1)
}

func sub() {
	atomic.AddInt32(&a, -1)
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second * 2)
	fmt.Printf("a: %v\n", a)
}
