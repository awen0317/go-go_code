package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f(){
	rand.Seed(time.Now().UnixNano())
	for i:=1;i<=5 ;i++  {
		r1:=rand.Int()
		r2:=rand.Intn(10)
		fmt.Println(r1,r2)
	}

}
func f1(i int)  {
	defer wg.Done()
	time.Sleep(time.Millisecond*time.Duration(rand.Intn(300)))
	fmt.Println(i)
}
var wg sync.WaitGroup
func main()  {
	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()
}