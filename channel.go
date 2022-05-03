package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

func sendValue() {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(10)
	fmt.Printf("send value: %v\n", value)
	time.Sleep(time.Second * 5)
	// 将数据放到channel中
	values <- value
}

func main() {
	// 函数结束时关闭channel
	defer close(values)
	go sendValue()

	fmt.Printf("\"wait\": %v\n", "wait")
	data := <-values // <-xxx 取出xxxchannel里的数据
	fmt.Printf("receive data: %v\n", data)
	fmt.Printf("\"end\": %v\n", "end")

	fmt.Printf("----------------\n")

	var ch = make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}

		close(ch)
	}()

	// for range遍历channel
	/* for v := range ch {
		fmt.Printf("v: %v\n", v)
	} */

	/* for {
		val, ok := <-ch
		if ok {
			fmt.Printf("val: %v\n", val)
		} else {
			// 不break会陷入死循环
			break
		}
	} */

	for i := 0; i < 3; i++ {
		v := <-ch
		fmt.Printf("v: %v\n", v)
	}
}
