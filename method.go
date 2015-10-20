package main

import (
	"fmt"
)

type myint int

func main() {
	var a myint
	a.increase()
	fmt.Println(a)
}

func (a *myint) increase() {
	*a = 100
	fmt.Println(a)
}
