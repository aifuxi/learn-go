package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		// 一次性读取文件全部内容
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "读取文件【%v】出错\n", filename)
			// 跳出当前循环
			continue
		}

		// 文件读出来是一个byte数组，需要转为string
		str := string(data)
		// fmt.Printf("data: %v\n", str)

		// 分割这个字符串，根据\n（换行符）
		for _, line := range strings.Split(str, "\n") {
			// fmt.Printf("第%d行: %v\n", n, line)
			// 把数据放到map里
			// 存入map，同时让这个字符串计数+1
			counts[line]++
		}
	}

	for key, value := range counts {
		fmt.Printf("字符串：【%v】，出现了%d次\n", key, value)
	}
}
