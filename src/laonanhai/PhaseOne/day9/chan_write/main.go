package main

import (
	"fmt"
	"time"
)

/**
写入的时候也会堵塞，用select也会使用结堵塞
 */
func main()  {
	ch :=make(chan int,1)
	go func() {
		var i int
		for true {
			select {
			case ch<-i:
			default:
				fmt.Println("channel is full")
				time.Sleep(time.Second)
			}
			i++
		}
	}()
	for true {
		v:=<-ch
		fmt.Println(v)
	}
}
