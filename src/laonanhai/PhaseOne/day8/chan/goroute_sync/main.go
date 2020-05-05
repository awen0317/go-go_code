package main

import (
	"fmt"
)

func main() {
	//var a int
	//a =10
	//fmt.Println(a)
	//
	//var s string
	//s :="abc"
	//fmt.Println(s)


	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	//for rang 可以检测close 然后关闭后不
	for v :=range ch{
		fmt.Println(v)
	}
}
