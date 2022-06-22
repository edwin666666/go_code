package main

import "fmt"

func main() {
	//根据值自行判定变量类型（类型推导）,编译器会按照变量赋予的值自行推导出变量类型，省略 int、string、bool等类型定义
	var i = 10.11
	fmt.Println("i=", i)
}
