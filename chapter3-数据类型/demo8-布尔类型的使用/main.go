package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//演示golang中bool类型使用
	var a = true
	var b = false
	fmt.Println("a=", a, "b=", b)
	//bool类型占用存储空间是1个字节
	fmt.Println("a和b占用的空间是", unsafe.Sizeof(a))
}
