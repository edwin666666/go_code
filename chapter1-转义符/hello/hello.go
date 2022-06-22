//开发第一个“hello”程序

//表示改hello.go文件所在的包是main，在 go 中， 每个文件都必须属于一个包。
package main

//表示：引用一个包，包名 fmt,引入该包后使用fmt包的函数。比如：fmt.Println
import "fmt"

//func是一个关键字，表示后边是一个函数
//main 是函数名，是一个主函数，即程序的入口
func main() {
	//表示调用 fmt 包的函数 Println 输出 "hello word"
	fmt.Println("hello word")
}
