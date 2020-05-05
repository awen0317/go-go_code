package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
		return
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0;i < flag.NArg(); i++{
		f,err :=os.Open(os.Args[0])
		if err !=nil{
			fmt.Fprintf(os.Stderr,"%s:error reading from %s,err:%s",os.Args[0],os.Args[i],err.Error())
			continue
		}
		//
		cat(bufio.NewReader(f))
	}
}
