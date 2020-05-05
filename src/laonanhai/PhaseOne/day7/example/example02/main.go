package main

import (
	"bufio"
	"fmt"
	"os"
)

type student struct {
	Name  string
	Age   int
	Score float32
}

var inputReader *bufio.Reader
var input string
var err error

func main() {
	var str = "stu01 18 300"
	var stu student
	//Sscanf的使用
	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)
	inputReader = bufio.NewReader(os.Stdin)
}
