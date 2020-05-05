package main

import (
	"fmt"
	"os"
)

func main()  {

	fmt.Printf("len %d\n",len(os.Args))
	for _,v :=range os.Args{
		fmt.Printf("len %s\n",v)
	}
}