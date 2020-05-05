package main

import (
	"fmt"
	"net"
)

func runServer(addr string)(err error)  {
	l,err :=net.Listen("tcp",addr)
	if err != nil {
		fmt.Println("listen failed",err)
	}

	for  {
		conn,err  :=l.Accept()
		if err != nil {
			fmt.Println("Accept failed",err)
			continue
		}
		go proccess(conn)
	} 
	
}
func proccess(conn net.Conn)  {
	defer  conn.Close()
	client :=&Client{
		conn: conn,
	}
	err :=client.Process()
	if err != nil {
		fmt.Println("clint process failed",err)
		return
	}
}
