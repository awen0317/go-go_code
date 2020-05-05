package main

import (
	"fmt"
)

func send(ch chan int,exit chan bool)  {
	for i:=0;i<10;i++{
		ch<-i
	}
	close(ch)
	exit<-true
}
func recv(ch chan int,exit chan bool)  {
	for {
		v, ok :=<-ch
		if !ok{
			break
		}
		fmt.Println(v)
	}
	exit<-true
}
func main() {
	ch :=make(chan int,10)
	exitChan :=make(chan bool,2)
	go send(ch,exitChan)
	go recv(ch,exitChan)

	for i:=0;i<2 ;i++  {
		<-exitChan
	}
	//time.Sleep(2*time.Second)
}
