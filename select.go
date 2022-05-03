package main

import (
	"fmt"
	"time"
)

var chInt = make(chan int)
var chStr = make(chan string)

func main() {
	go func() {
		chInt <- 100
		chStr <- "heihei"

		close(chInt)
		close(chStr)
	}()

	for {
		select {
		case r := <-chInt:
			fmt.Printf("chInt: %v\n", r)
		case r := <-chStr:
			fmt.Printf("chStr: %v\n", r)
		default:
			fmt.Printf("default\n")
		}
		time.Sleep(time.Second)
	}
}
