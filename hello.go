package main

import "fmt"

const (
	B int = 1 << (iota * 10)
	KB
	MB
	GB
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

	if a := 1; a > 0 {
		fmt.Println(a)
	}

	a := 1
	for {
		if a > 3 {
			fmt.Println(a)
			break
		}
		a++
	}
	fmt.Println(a)
}
