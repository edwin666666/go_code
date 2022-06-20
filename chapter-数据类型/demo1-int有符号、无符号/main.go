package main

import "fmt"

//演示golang中整数类型使用
func main() {
	//有符号(负数、正数) int8 的范围，int16、int32、int64 以此类推
	var b int8 = -128
	var c int8 = 127
	fmt.Println("b=", b, "c=", c)
	//无符号（正数） uint8 的范围，int16、int32、int64 以此类推
	var d uint8 = 0
	var e uint8 = 255
	var f uint16 = 65535
	var g uint32 = 4294967295
	var h uint64 = 18446744073709551615
	fmt.Println("d=", d, "e=", e, "f=", f, "g=", g, "h=", h)

}
