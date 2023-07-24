package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// 1. 先获取输入的文件名
	for _, filename := range os.Args[1:] {
		// fmt.Printf("filename: %v\n", filename)
		counts := make(map[string]int)

		// 读取文件
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
			counts[line]++
		}

		// 遍历map，如果存在value > 1，说明这个文件存在重复行
		for _, v := range counts {
			if v > 1 {
				fmt.Printf("文件 【%v】存在重复行\n", filename)
				break
			}
		}
	}
}
