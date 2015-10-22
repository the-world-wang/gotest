package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go testGo(c)
	b := <-c
	fmt.Println(b)
}

func testGo(c chan bool) {
	fmt.Println("Go Go Go")
	c <- true
}
