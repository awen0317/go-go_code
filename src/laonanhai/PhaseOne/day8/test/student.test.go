package main

import (
	"testing"
)


func TestSave(t *testing.T)  {
	stu :=&student{
		ame: "stu01",
		Sex: "man",
		Age: 10,
	}
	err :=stu.Save()
	if err !=nil{
		t.Fatalf("save faild %v",err)
	}

}
func TestLoad(t *testing.T)  {
		stu :=&student{
			ame: "stu01",
			Sex: "man",
			Age: 10,
		}
		err :=stu.Save()
	if err !=nil{
		t.Fatalf("save faild %v",err)
	}
		stu02 :=&student{}
		err = stu02.Load()
	if err !=nil{
		t.Fatalf("load faild %v",err)
	}
	if stu.ame!=stu02.ame{
		t.Fatalf("load faild %v, name not equal",err)
	}
	if stu.Sex!=stu02.Sex{
		t.Fatalf("load faild %v, name not equal",err)
	}

}

