package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func printErr(err error)  {
	if err !=nil{
		fmt.Println("read string failed,err",err)
		return
	}
}

type CharCount struct {
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}

func main()  {

	var count CharCount
	file,err := os.Open("H:/test.txt")
	printErr(err)
	defer file.Close()
	reaader := bufio.NewReader(file)

	for {
		str,err :=reaader.ReadString('\n')
		if err == io.EOF && len(str)==0{
			break
		}
		if err !=nil{
			fmt.Println("read string failed,err",err)
			break
		}
		runeStr :=[]rune(str)
		for _,v :=range runeStr{
			switch  {
			case v>='a' && v <='z':
				fallthrough
			case v>='A' && v <='Z':
				count.ChCount++
			case v==' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++

			}
		}

	}
	fmt.Println(count)



}