package main

import (
	"fmt"
)

type Test interface {
	Print()
}

type Student struct {
	name string
	age int
	score float64
}

func (p *Student) Print()  {
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.score)
}

func main()  {
	var t Test
	var stu Student =Student{
		name:  "stu1",
		age:   19,
		score: 99.1,
	}
	t = &stu
	t.Print()
}