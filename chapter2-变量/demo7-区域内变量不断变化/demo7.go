package main

import "fmt"

//该区域的数值可以在同一类型范围内不断变化（重点）
func main() {
	var i int = 10
	i = 20
	i = 30
	fmt.Print("i=", i)
}
