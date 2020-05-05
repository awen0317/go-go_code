package main

import (
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

func CopyFile(dsName string,srcName string)(written int64,err error)  {
	src ,err :=	os.Open(srcName)
	printErr(err)
	defer src.Close()
	dst,err :=os.OpenFile(dsName,os.O_WRONLY|os.O_CREATE,0664)
	printErr(err)
	defer dst.Close()
	return io.Copy(dst,src)


}
func main()  {

	CopyFile("target.txt","source.txt");
	fmt.Println("Copy Done")

}