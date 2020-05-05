package main

import (
	"fmt"
	"time"
)


type Car struct {
	Name string
	Age int
}
type Car2 struct {
	Name string


}

type Train struct {
	Car
	Car2
	createTime time.Time
	int
}

func (t *Train) Run(age int) {
	t.int = age
}

func main()  {

	var train Train
	fmt.Println(train)
	//train.int = 300
	//train.Run(100)
	//
	//train.Car.Name ="test"
	//fmt.Println(train.int)
}
