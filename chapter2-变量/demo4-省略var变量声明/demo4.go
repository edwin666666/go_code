package main

import "fmt"

func main() {
	//省略var，注意 := 左侧的变两个不应噶爱是已经声明过的，否则会导致编译错误
	//下面等价 var name string = "侯利军"
	//注意 := , 冒号是不能省略，否则报错
	name := "侯利军"
	fmt.Println("NAME=", name)
}
