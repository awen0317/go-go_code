package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json，其他语言，够语言能够识别的结构题

type Student struct {
	Name string
	Age int
}
type Student2 struct {
	Name string
	Age int
	Sex string
}
func main()  {
	p1 :=&Student{
		Name: "a",
		Age:  12,
	}
	b,err :=json.Marshal(p1)
	if err != nil {
		fmt.Println("json 序列化成功")
	}
	fmt.Println(string(b))

	c :=`{"Name":"a","Age":13}`
	p2 :=&Student2{}
	err = json.Unmarshal([]byte(c),&p2)
	fmt.Println(p2.Name)
	fmt.Println(p2.Age)
}
