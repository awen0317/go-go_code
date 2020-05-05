package main

import (
	"fmt"
	"time"
)

func write(ch chan int)  {
	for i:=0;i<100;i++{
		ch <-i
	}
}
func read(ch chan int)  {
	for  {
		var b int
		b= <-ch
		fmt.Println(b)
	}
}
func main()  {
	intChin := make(chan int,10)
	go write(intChin)
	go read(intChin)
	time.Sleep(10*time.Second)
}
