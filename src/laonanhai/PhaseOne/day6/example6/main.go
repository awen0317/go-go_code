package main

import (
	"fmt"
)

func just(items ...interface{})  {
	fmt.Printf("%T",items)
	fmt.Println(items)
	for _,v :=range items{
		switch v.(type) {
		case bool:
			fmt.Println("bool")
		case int,int64,int32:
			fmt.Println("int64")
		case float32:
			fmt.Println("float32")
		case float64:
			fmt.Println("float64")
		case string:
			fmt.Println("string")

		}
	}
}

func main()  {

	//var a interface{}
	//var b int
	//a =b
	//c :=a.(int)
	//fmt.Printf("%T",c)
	just(8.2,1,"abc")
}