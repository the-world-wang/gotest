package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=myfirst sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	//增
	stmt, _ := db.Prepare("insert into student(id,age,name) values ($1,$2,$3)")
	stmt.Exec(4, 100, "gaofushuai")

	//删
	id := 1
	db.Exec("delete from student where id=$1", id)

	//查
	db.Exec("update student set name=$1", "王浩")

	/*查看结果*/
	rows, _ := db.Query("select * from student")
	for rows.Next() {
		var id int
		var age int
		var name string
		rows.Scan(&id, &age, &name)
		fmt.Printf("id is %d,age is %d,name is %s\n", id, age, name)
	}
}
