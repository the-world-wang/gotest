package main

import (
	"fmt"
	"sort"
)

func main() {
	// map的声明
	var m map[string]string //[]内为key的类型，外面为value的类型

	// map的创建
	// 1.make
	m = make(map[string]string)
	// 2.new
	mm := new(map[int]string)
	fmt.Println(mm)
	//(*mm)[1] = "name" , 编译不通过
	//(*mm)[2] = "age"
	//fmt.Println(mm)

	// 3.直接原生

	// map的赋值
	m["1"] = "一"
	m["2"] = "二"
	fmt.Println(m, m["1"])

	// map的删除
	delete(m, "1")
	fmt.Println(m)

	v, ok := m["3"]
	if !ok {
		fmt.Println("m[3]不存在=", v)
	}

	// 排序map的key
	m2 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "f"}
	sm := make([]int, 5) // 预定义空间
	i := 0
	for k, _ := range m2 {
		sm[i] = k
		i++
	}

	fmt.Println(sm)
	sort.Ints(sm)
	fmt.Println(sm)

	m3 := make(map[string]int)

	// 使用for range对map进行遍历
	for k, v := range m2 {
		m3[v] = k
	}
	fmt.Println(m3)

}
