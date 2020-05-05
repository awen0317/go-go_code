package main

import (
	"encoding/json"
	"fmt"
)


func printErr(err error)  {
	if err !=nil{
		fmt.Println("read string failed,err",err)
		return
	}
}
type User struct {
	UserName string
	NickName string
	Age int
	Sex string
	Email string
}

func testStruct()  {
	use1 := &User{
		UserName: "awen",
		NickName: "a",
		Age:      10,
		Sex:      "男",
		Email:    "714833301@qq.com",
	}
	data,err :=json.Marshal(use1)
	printErr(err)
	fmt.Printf("%s",string(data))

}
func testInt()  {
	var a int =100
	data,err :=json.Marshal(a)
	printErr(err)
	fmt.Printf("%s",string(data))

}
func testMap()  {
	m := make(map[string]interface{})
	m["name"] = "yutiawne"
	m["age"] = 18
	data,err :=json.Marshal(m)
	printErr(err)
	fmt.Printf("%s",string(data))
}
func testSlice()  {
	m := make(map[string]interface{})
	var s []map[string]interface{}
	m["name"] = "yutiawne"
	m["age"] = 18
	s =append(s,m)
	m = make(map[string]interface{})
	m["name"] = "12345"
	m["age"] = 17
	s =append(s,m)
	data,err :=json.Marshal(s)
	printErr(err)
	fmt.Printf("%s",string(data))




}
func main()  {
	//testStruct()
	//testInt()
	//testMap()
	testSlice()



}