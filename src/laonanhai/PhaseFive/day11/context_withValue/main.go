package main

import (
	"fmt"
	"golang.org/x/net/context"
	"sync"
	"time"
)

type TraceCode string

var wg sync.WaitGroup

func work(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("work trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond*10)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Println("work done")
	wg.Done()
}
func main() {
	// 设置一个50秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	// 在系统的入口中设置trace code传递给后续的启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "151531531")
	wg.Add(1)
	go work(ctx)
	time.Sleep(time.Second * 5)
	cancel()

}
