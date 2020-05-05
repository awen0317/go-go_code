package main

import (
	"fmt"
	"golang.org/x/net/context"
	"sync"
	"time"
)
/**
主通知子协程结束
 */
var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f2(ctx context.Context)  {
	defer wg.Done()
FORLOOP:
	for true {
		fmt.Println("zhoubaocen")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:

		}
	}

}
func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
FORLOOP:
	for true {
		fmt.Println("zhoulin")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:

		}
	}
}
func main() {
	ctx,cancel :=context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
