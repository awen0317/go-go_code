package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m=make(map[int]uint)
	lock sync.Mutex
)

type Task struct {
	n int
}

func factorial(t *Task) {
	var sum uint = 1

	for i:=1;i<=t.n;i++{
		sum *=uint(i)
	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}
func main() {
	var t *Task
	for i := 1; i <= 56; i++ {
		t = &Task{n: i}
		go factorial(t)
	}
	time.Sleep(3 * time.Second)

	fmt.Println(m)
}
