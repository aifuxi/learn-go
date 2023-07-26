package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		r2, _, err := r.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("读取错误%v", err)
		}

		fmt.Printf("你输入的是: %v\n", string(r2))
	}
}
