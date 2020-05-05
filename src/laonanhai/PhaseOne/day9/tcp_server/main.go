package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn)  {

	defer conn.Close()
	for {
		buf :=make([]byte,512)
		_,err :=conn.Read(buf)
		if err!=nil{
			fmt.Println("read err",err)
			return
		}
		fmt.Println("read ",string(buf))

	}
}
func main()  {

	fmt.Println("server start")
	listen,err :=net.Listen("tcp_server","0.0.0.0:50000")
	if err!=nil{
		fmt.Println("listen failed err",err)
		return
	}
	for true {
		conn,err :=listen.Accept()
		if err!=nil{
			fmt.Println("accept failed err",err)
			return
		}
		go process(conn)

	}
}
