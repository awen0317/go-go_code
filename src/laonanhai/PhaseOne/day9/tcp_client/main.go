package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func process(conn net.Conn)  {

}
func main()  {
	conn,err :=net.Dial("tcp","localhost:50000")
	if err !=nil{
		fmt.Println("Error Dialing",err.Error())
	}
	defer  conn.Close()
	inputReader :=bufio.NewReader(os.Stdin)
	for {
		input,_ := inputReader.ReadString('\n')
		trimmedinput :=strings.Trim(input,"\r\n")
		if trimmedinput=="Q"{
			return
		}
		_,err =conn.Write([]byte(trimmedinput))
		if err !=nil{
			return
		}
	}


}
