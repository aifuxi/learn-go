package main

import "fmt"

// 接口一般以er结尾
type USBer interface {
	read()
	write()
}

type Computer struct {
	brand string
}

func (c Computer) read() {
	fmt.Printf("c.brand: read %v\n", c.brand)
}

func (c Computer) write() {
	fmt.Printf("c.brand: write %v\n", c.brand)
}

type Phone struct {
	speed string
}

func (p Phone) read() {
	fmt.Printf("p.speed: read %v\n", p.speed)
}

func (p Phone) write() {
	fmt.Printf("p.speed: write %v\n", p.speed)
}

func main() {
	c1 := Computer{
		brand: "msi",
	}
	c1.read()
	c1.write()

	p1 := Phone{
		speed: "4g",
	}
	p1.read()
	p1.write()
}
