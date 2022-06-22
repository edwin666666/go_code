package main

import "fmt"

//定义全局变量,方法1
var n1, n2, name = 100, 200, "侯利军"

//定义全局变量，方法2
var (
	n3    = 1
	n4    = 2
	name2 = "hello"
)

func main() {
	fmt.Println("n1=", n1, "n2=", n2, "name=", name)
	fmt.Println("n3=", n3, "n4=", n4, "name2=", name2)
}
