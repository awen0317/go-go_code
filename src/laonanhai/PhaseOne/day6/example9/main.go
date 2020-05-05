package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	Sex   string
}

func (s Student) Set(name string,age int,score float32,sex string) {
	s.Name=name
	s.Age=age
	s.Score=score
	s.Sex=sex
}

func (s Student) Print()  {
	fmt.Println("-----start----")
	fmt.Println(s)
	fmt.Println("-----end----")
}

func TestStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("except struct")
		return
	}
	num := val.NumField()
	fmt.Printf("STRUCT HAS %d fields \n", num)
	num1 := val.NumMethod()
	fmt.Printf("STRUCT HAS %d memthod \n", num1)
	var params []reflect.Value
	val.Method(0).Call(params)

}
func main() {
	var a Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 92.8,
		Sex:   "",
	}
	TestStruct(a)
	fmt.Println(a)
}
