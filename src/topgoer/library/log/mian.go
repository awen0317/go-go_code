package main

import (
	"log"
	"os"
)

func main() {
	//logFile,err := os.OpenFile("./xx.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	//if err !=nil{
	//	panic(err)
	//}
	//log.SetOutput(logFile)
	//log.SetFlags(log.Llongfile|log.Lmicroseconds|log.LUTC|log.Ldate)
	//log.Println("这是一条很普通的日志")
	//log.SetPrefix("【oncall系统】")
	//log.Println("这是一条很普通的日志")
	//v :="很普通"
	//log.Printf("这是一条%s的日志",v)
	//log.Fatalln("这一条会触发fatal,日志")
	//log.Panicln("这一条会触发panic,日志")
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}
