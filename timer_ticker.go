package main

import "time"

func main() {
	println("开始", time.Now().String())
	// timer只会执行一次
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	println("结束", time.Now().String())

	// ticker会周期性的运行
	ticker1 := time.NewTicker(time.Second * 2)
	sum := 0
	for v := range ticker1.C {
		if sum >= 5 {
			ticker1.Stop()
			break
		}
		sum += 1
		println("跳", v.String())
		println("sum", sum)
	}
}
