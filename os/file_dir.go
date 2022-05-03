package main

import (
	"fmt"
	"os"
)

func createFile() {
	f, err := os.Create("test_create_file.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("f: %v\n", f.Name())
}

func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("dir: %v\n", dir)
}

func mkDir() {
	err := os.Mkdir("test_os_mkdir", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err2 := os.MkdirAll("test_os_mkdir_all/a/b/c", os.ModePerm)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

func renameFile() {
	err := os.Rename("test_create_file.txt", "test_create_file_renamed.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func readFile() {
	b, err := os.ReadFile("test_write_file.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("b: %v\n", string(b))
}

func writeFile() {
	err := os.WriteFile("test_write_file.txt", []byte("hello go"), os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func main() {
	// createFile()
	// getWd()
	// mkDir()
	// renameFile()
	// writeFile()
	readFile()
}
