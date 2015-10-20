package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
}

func main() {
	user := User{1, "wanghao"}
	Info(user)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println(t.Name())

	vaule := reflect.ValueOf(o)
	for i := 0; i < vaule.NumField(); i++ {
		field := vaule.Field(i)
		fmt.Println(field.Type())
	}
}
