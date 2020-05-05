package main

import (
	"fmt"
	"runtime"
)

func test()  {
	defer func() {
		if err :=recover();err !=nil{
			fmt.Println("panic",err)
		}
	}()
	var m map[string]string
	m["stu"] = "100"
}
func main()  {
	num :=runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	for i:=0;i<100 ;i++  {
		go test()
	}

}
