package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notify bool

// func f(){
// 	defer wg.Done()
//
// 	for true {
// 		fmt.Println("周林")
// 		time.Sleep(time.Millisecond*500)
// 		if notify{
// 			break
// 		}
// 	}
// }
func f(i int){
	wg.Add(1)
	defer wg.Done()
	for true {
		fmt.Println("周林"+strconv.Itoa(i))
		time.Sleep(time.Millisecond*500)
		if notify{
			break
		}
	}
}
func main() {
	for i:=1;i<=5 ;i++  {
		go f(i)
	}
	time.Sleep(time.Second*10)
	notify =true
	wg.Wait()
}
