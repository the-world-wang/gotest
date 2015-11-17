package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var (
	engine *xorm.Engine
)

type student struct {
	Id   int64
	Age  int64
	Name string
}
type person struct {
	Id   int
	Age  int
	Name string
}
type personAndStudent struct {
	Id   int `xorm:"IDL"`
	Age  int
	Name string
}

func main() {
	var err error
	engine, err = xorm.NewEngine("postgres", "user=postgres password=123456 dbname=myfirst sslmode=disable")
	if err != nil {
		panic(err)
	}

	student1 := new(student)
	session := engine.NewSession()
	session.Id(2).Get(student1)
	fmt.Println(student1)
	student1.Name = "wohao10"
	student1.Age = 100

	session.Id(3).Cols("name", "age").Update(student1)
}
