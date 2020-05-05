package main

import (
	"bufio"
	"fmt"
	"os"
)

func printErr(err error)  {
	if err !=nil{
		fmt.Println("read string failed,err",err)
		return
	}
}

func main()  {
	//只读的话用open，要写的话用openfile
	file,err :=os.Open("H:/test.log.txt")
	printErr(err)
	defer file.Close()
	reader :=bufio.NewReader(file)
	str,err :=reader.ReadString('\n')
	printErr(err)

	fmt.Printf("read str ucc,ret:%s\n",str)
}