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
func testMap() (ret string,err error) {
	m := make(map[string]interface{})
	m["name"] = "yutiawne"
	m["age"] = 18
	data,err := json.Marshal(m)
	if err != nil {
		err =fmt.Errorf("json.Marsh1 failded ,err ",err)
	}
	ret =string(data)
	fmt.Printf("%s",string(data))
	return
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
func testStructUnMarsh() (ret string,err error) {
	use1 := &User{
		UserName: "awen",
		NickName: "a",
		Age:      10,
		Sex:      "男",
		Email:    "714833301@qq.com",
	}
	data,err :=json.Marshal(use1)
	if err != nil {
		err =fmt.Errorf("json.Marsh1 failded ,err ",err)
	}
	fmt.Printf("%s",string(data))
	ret =string(data)
	return
}
func test()  {
	data,err :=testStructUnMarsh()
	if err != nil {
		fmt.Println("test struct failed ")
	}
	var user1 User
	json.Unmarshal([]byte(data),&user1)
	if err != nil {
		fmt.Println("错误信息")
		return
	}
	fmt.Println(user1)
}
func test2()  {
	data,err := testMap()
	if err != nil {
		fmt.Println("test struct failed ")
	}
	m :=make(map[string]interface{})
	json.Unmarshal([]byte(data),&m)
	if err != nil {
		fmt.Println("错误信息")
		return
	}
	fmt.Println(m)
}
func main()  {
	testStruct()
	test2()
	//testInt()
	//testMap()
	//testSlice()



}