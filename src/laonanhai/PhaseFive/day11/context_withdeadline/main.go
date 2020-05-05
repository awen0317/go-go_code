package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func main()  {
	d :=time.Now().Add(2000*time.Millisecond)
	ctx,cancel :=context.WithDeadline(context.Background(),d)
	defer cancel()
	select {
	case <-time.After(time.Second):
		fmt.Println("zhoulin")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
