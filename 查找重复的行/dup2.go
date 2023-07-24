package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 先声明一个map，用来存储出现的字符串，value是出现的次数
	counts := make(map[string]int)

	// 读取运行时参数
	files := os.Args[1:]

	fmt.Printf("files: %v\n", files)

	if len(files) == 0 {
		fmt.Println("没有文件")
	} else {
		// 先把files输出看看
		for _, file := range files {
			// fmt.Printf("v: %v\n", v)

			// 打开文件, // !: 注意：文件路径是以当前执行命令的路径算的
			f, err := os.Open(file)

			// 读取文件出错
			if err != nil {
				fmt.Fprintf(os.Stderr, "读取文件【%v】出错\n", file)
				// 跳出当前循环
				continue
			}

			// 读取文件正常
			countLines(f, counts)
		}
	}

	// 输出counts %+v会输出map的key
	// fmt.Printf("counts: %+v\n", counts)
	// !: 注意：map的遍历不是有序的
	for key, value := range counts {
		fmt.Printf("字符串：【%v】，出现了%d次\n", key, value)
	}
}

func countLines(f *os.File, counts map[string]int) {
	// 从文件中读取数据
	s := bufio.NewScanner(f)
	for s.Scan() {
		// 将读取到的数据转为文本
		s2 := s.Text()
		// 存入map，同时让这个字符串计数+1
		counts[s2]++
	}
}
