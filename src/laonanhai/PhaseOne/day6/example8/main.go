package main

import (
	"fmt"
	"reflect"
)

func test(b interface{}){
	t :=reflect.TypeOf(b)
	fmt.Println(t)
	v :=reflect.ValueOf(b)
	 k :=v.Kind()
	 fmt.Println(k)
}
func main()  {
	var a int =100
	test(a)
}