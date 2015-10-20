package main

import (
	"fmt"
)

const max int = 512
const min = 42.5

//预定于常量iota
// 1.iota是常量计数器，每定义一个常量，iota增加1
// 2.没遇到一个const，iota重置为0
// 3.可以用来枚举
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 全局变量可以通过var()来声明，并且可以并行
// 局部变量不可以使用var()来声明，但是可以并行
var (
	name        string = "wanghao"
	age, height        = 12, 170
)

func main() {
	var a int = 100
	var b = 100
	c := int(100)
	d := 100

	//可以并行
	var e, f, g int = 1, 2, 3

	fmt.Println(a, b, c, d, max, min, name, age, height, e, f, g, Monday)
}
