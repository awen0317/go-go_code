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
	reader :=bufio.NewReader(os.Stdin)
	str,err :=reader.ReadString('\n')
	printErr(err)

	fmt.Printf("read str ucc,ret:%s\n",str)
}