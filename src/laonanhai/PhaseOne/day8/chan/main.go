package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func main() {
	//var intChan chan int
	//intChan = make(chan int,10)
	//intChan<-10
	//var mapChan chan map[string]string
	//mapChan = make(chan map[string]string,10)
	//m :=make(map[string]string,16)
	//m["stu01"] = "stu1"
	//m["stu02"] = "stu2"
	//mapChan<-m

	//var structChan chan *Student
	//structChan = make(chan *Student,10)
	//stu := &Student{Name: "xiaoming"}
	//
	//structChan<- stu
	var structChan chan interface{}
	structChan = make(chan interface{},10)
	//
	stu := &Student{Name:"zhangsan"}

	structChan<-stu

	var stu01 interface{}
	var stu02 *Student
	stu01 = <-structChan
	if stu02,ok := stu01.(*Student);!ok{
		fmt.Println("转移失败")

	}
	fmt.Println(stu02)
}
