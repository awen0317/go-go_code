package main

import (
	"fmt"
	"time"
)

func main()  {
	for i:=0;i<1024;i++{
		go func() {
			for  {
				t :=time.NewTicker(time.Second)
				select {
				case <-t.C:
					fmt.Println("chaoshi le ")
					
				}
				t.Stop()
			}
		}()
	}
}
