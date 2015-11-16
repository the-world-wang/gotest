package main

import (
	"fmt"
	"log"
)

func main() {
	//log.Fatal("hello")
	log.SetOutput()
	log.SetFlags(flag)

	log.New(out, "[]", flag)

	// log其实核心还是Logger,可以通过New(out,prefix,flag)来创建
	// 默认会建一个标准的输入
}
