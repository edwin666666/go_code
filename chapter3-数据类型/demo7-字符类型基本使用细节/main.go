package main

import (
	"fmt"
)

func main() {
	var a byte = 'a'
	var b byte = '0'
	//当我们直接输出byte值，就是输出了对应的阿斯克码（ASCII）的码值
	fmt.Println("a=", a, "\nb=", b)
	//如果我们希望输出字符，需要使用格式化输出
	fmt.Printf("a=%c \nb=%c", a, b)
	var c int = '北'
	fmt.Printf("\nc=%c\n", c)
	//字符类型是可以进行运算，相当于一个证书，运算时是按照码值运行
	var d = 10 + 'a' // 10+97=107
	fmt.Println("d=", d)
}
