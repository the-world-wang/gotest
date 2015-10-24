package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Id   int    `json:"id"` //不是单引号
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func main() {
	p := person{1, 22, "wanghao"}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error is", err)
	}
	fmt.Println(string(b))

	var p2 person
	json.Unmarshal(b, &p2) // 在已知解析类型的情况下，使用这种方式
	fmt.Println(p2)

	var p3 interface{}
	json.Unmarshal(b, &p3)
	m := p3.(map[string]interface{})
	for k, v := range m {

	}

}
