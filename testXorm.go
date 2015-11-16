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
	fmt.Println(engine)

	// engine.Table("studentAndTeacher")
	//s := student{7, 22, "wanghao"}
	/*engine.Table("student")
	engine.Sync(s)*/

	/*	dbm, _ := engine.DBMetas()
		fmt.Println(dbm)*/

	/*_, err = engine.Insert(s)
	if err != nil {
		fmt.Println(err)
	}*/

	/*isTable, _ := engine.IsTableExist("teacher")
	fmt.Println(isTable)

	fmt.Println(engine.IsTableEmpty("teacher"))

	var ss student
	engine.Id(6).Get(&ss)
	fmt.Println(ss)

	s := student{2, 2}
	engine.Sql("select * from student", &s)*/

	/*获取表的结构*/
	table, _ := engine.DBMetas()
	for i, v := range table {
		fmt.Printf("i is %s , v is %s", i, v)
	}

	/*创建表*/
	engine.CreateTables(new(person))
	b, _ := engine.IsTableExist("person")
	fmt.Println("is the table person exists", b)

	/*删除表*/
	/*	engine.DropTables(new(person))
		boo, _ := engine.IsTableExist("person")
		fmt.Println("is the table person exists", boo)*/

	/*插入一条数据*/
	//engine.Insert(person{2, 22, "wanghao"})

	p := person{Id: 2}
	engine.Get(&p)
	fmt.Println(p)

	p2 := new(person)
	engine.Where("id = $1", 2).Get(p2)
	fmt.Println("p2 is..", p2)

	persons := make([]person, 0)
	engine.Find(&persons)
	fmt.Println(persons)

	/*personMap := make(map[int64]person)
	engine.Find(&personMap)*/
	// session := engine.NewSession()
	pas := personAndStudent{1, 22, "wanghao"}
	engine.CreateTables(&pas) //生成对应表的结构为person_and_student

}
