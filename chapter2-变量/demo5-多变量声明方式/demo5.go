package main

import "fmt"

// func main() {
// 	//一次性声明多个变量,方式1
// 	var n1, n2, n3 int
// 	fmt.Println(n1, n2, n3)
// }

// func main() {
// 	//一次性声明多个变量,方式2
// 	var n1, name, n3 = 100, "tom", 888
// 	fmt.Println(n1, name, n3)
// }

func main() {
	//一次性声明多个变量,类型推导，方式3
	n1, name, n3 := 100, "侯利军", 888
	fmt.Println(n1, name, n3)
}
