package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//查看变量占用字节大小和数据类型
	var a byte = 12
	//unsafe.Sizeof(a)是unsafe包的一个函数，可以返回a变量占用的字节数
	fmt.Printf("a 的数据类型 %T\na 占用的字节大小 %d", a, unsafe.Sizeof(a))
}
