package main

import (
	"fmt"
)

//int、uint、byte(0-255)、rune使用(视频中没有测试)，使用默认值跟主机操作系统有关联，操作系统64位8个字节，操作系统32位4个字节
func main() {
	//最小
	var a int = -9223372036854775808
	//最大
	var b int = 9223372036854775807
	//最小
	var c uint = 0
	//最大
	var d uint = 18446744073709551615
	var e byte = 0
	var f byte = 255
	fmt.Println("int最小值\n", "a=", a, "\nint最大值", "\nb=", b, "\nuint最小值", "\nc=", c, "\nuint最大值", "\nd=", d, "\ne=", e, "\nf=", f)
}
