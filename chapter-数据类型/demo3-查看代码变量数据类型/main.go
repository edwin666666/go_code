package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//整型使用的细节
	//查看代码中变量的数据类型
	// var a = 100
	// fmt.Printf("a 的类型 %T", a)

	//查看代码中变量占用的字节大小和数据类型
	var b = 100
	//unsafe.Sizeof(b),是unsafe包的一个函数，可以返回b变量占用的字节数
	fmt.Printf("b 的占用的字节大小 %d  b 的数据类型 %T", b, unsafe.Sizeof(b))
}
