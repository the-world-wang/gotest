package main

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
)

func main() {
	fmt.Println(uuid.NewRandom())
	fmt.Println(uuid.New())
}
