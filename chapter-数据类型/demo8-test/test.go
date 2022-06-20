package main

import (
	"fmt"
)

func main() {
	var a byte = 'a'
	var b byte = '0'
	//当我们直接输出byte值，就是输出了对应的阿斯克码（ASCII）的码值
	fmt.Println("a=", a, "\nb=", b)
}
