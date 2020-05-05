package main

import (
	"fmt"
	"time"
)

func suShu(entrych chan int, resch chan int,countchan chan bool) {
	for v := range entrych {
		b := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				b =false
			}
		}
		if b{
			resch<-v
		}
	}
	countchan<-true

}
func main() {
	entryChan := make(chan int, 1000)
	resChan := make(chan int, 1000)
	countChan := make(chan bool, 8)

	for i := 1; i <= 1000; i++ {
		entryChan <- i
	}
	close(entryChan)
	time.Sleep(5*time.Second)
	for i := 0; i < 8; i++ {
		go suShu(entryChan, resChan,countChan)
	}
	//for i := 0; i < 8; i++ {
	//	<-countChan
	//}
	//close(resChan)

	go func() {
		for i := 0; i < 8; i++ {
			<-countChan
		}
		close(resChan)
	}()

	for v :=range resChan{
		fmt.Println(v)
	}
}
