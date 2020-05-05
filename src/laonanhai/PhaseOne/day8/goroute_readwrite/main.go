package main

import (
	"fmt"
)

var a int
func main() {
	//time.Sleep(2*time.Second)
	ch := make(chan int, 10)
	for i := 1; i < 10; i++ {
		ch <- i
	}
	close(ch)

	for {
		if a>10{
			break
		}
		//var b int
		//b = <-ch
		//fmt.Println(b)

		select {
		case v:=<-ch:
			fmt.Println(v)
		default:
			fmt.Println("game over")
		}
		a++
	}
}
