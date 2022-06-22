package main

import (
	"fmt"
)

func main() {
	//小数类型也叫浮点数类型
	var price float32 = 13.59
	fmt.Println("price=", price)
	var a float32 = -0.0000012345
	var b float64 = -123456789.987654321
	fmt.Println("a=", a, "\nb=", b)
	//尾数部分可能丢失，造成精度损失。-123.0000901,比较 c 和 d 那个精度更加准确
	var c float32 = -123.0000901
	var d float64 = -123.0000901
	fmt.Println("c=", c, "\nd=", d)
	//查看不声明float32或者float64,Golang默认使用float64
	var e = 1.111
	fmt.Printf("\ne的数据类型是 %T", e)
	//十进制数形式：如：5.12	.512（必须有小数点）
	f := 5.12
	g := .512
	fmt.Println("\nf=", f, "\ng=", g)
	//科学计数法形式，5.1234e2 等价于 5.1234 乘以 10的2次方（*100,512.34）,小写 e 和 大写 E 一样
	h := 5.1234e2
	i := 5.1234e2
	// e-2 等价于 5.1234 除以 10的2次方（除以100,0.051234）
	j := 5.1234e-2
	fmt.Println("h=", h, "\ni=", i, "\nj=", j)
}
