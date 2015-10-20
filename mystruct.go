package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type student struct {
	person
	School string
}

type teacher struct {
	person
	Dept string
}

func main() {

	a := person{}
	a.Name = "wanghao"
	a.Age = 22

	fmt.Println(a)

	//利用聚合完成继承
	s := student{School: "yangzhou", person: person{Name: "wanghao", Age: 22}}
	t := teacher{Dept: "研发", person: person{Name: "xudan", Age: 35}}

	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(s.Name)
}
