package main

import (
	"fmt"
)

func main() {
	var a [2]int
	fmt.Println(a)

	b := [2]int{1, 2}
	fmt.Println(b)

	// new 创建对象
	p := new([10]int)
	fmt.Println(*p)

	// 多维数组
	c := [3][4]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4}}
	fmt.Println(c)

	// 支持for range用于遍历
	for i, v := range b {
		fmt.Println("array", i, "=", v)
	}

	// 数组是值类型，每次传递产生一个副本
	// 可以看出并没有修改数组的值
	Change(b)
	for i, v := range b {
		fmt.Println("array", i, "=", v)
	}
}

func Change(b [2]int) {
	b[0] = 3
}
