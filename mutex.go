package main

import (
	"fmt"
	"sync"
)

var num = 100
var wp sync.WaitGroup

// 互斥锁
var lock sync.Mutex

func add() {
	defer wp.Done()
	defer lock.Unlock()

	lock.Lock()
	num += 1
	fmt.Printf("num++: %v\n", num)
}

func sub() {
	defer wp.Done()
	defer lock.Unlock()

	lock.Lock()
	num -= 1
	fmt.Printf("num--: %v\n", num)
}

func main() {
	for i := 0; i < 100; i++ {
		wp.Add(1)
		go add()

		wp.Add(1)
		go sub()
	}

	// Wait只是保证协程不会在main结束后被终止，并不能保证数据的一致性
	wp.Wait()

	fmt.Printf("end num: %v\n", num)
}
