package main

import (
	"fmt"
)

func main() {

	// slice通过指针，利用数组，实现变长的数组
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 切片slice的第一种构造方式
	s1 := a[5:]
	s2 := a[:5]
	s2[0] = 10
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(a)

	// slice类似与java中的list，初期创建可以指定一个限制，但是当这个限制不够的时候，就会重新分配内存
	s3 := make([]int, 3, 10)
	s4 := make([]int, 3)

	// 元素的个数和容量的空间是两个不同的值
	fmt.Println(len(s3), cap(s3))
	fmt.Println(len(s4), cap(s4))

	// 当增加到超过切片的cap，就会重新为切片分配内存空间
	s3 = append(s3, 1, 2, 3, 4, 5, 6, 7)
	for i, v := range s3 {
		fmt.Println(i, "=", v)
	}

	s5 := s3[:]
	fmt.Println(s5)

	//make和new的比较
	//new返回的指向这个切片的指针，相当于指针的指针
	//而make是直接将生成的对象返回，并且可以指定初始化的cap
	s6 := new([]int)           // 打印出来是	&[]
	*s6 = append(*s6, 1, 2, 3) // 同样也可以利用指针append
	fmt.Println(s6)            // 	&[1 2 3]
}
