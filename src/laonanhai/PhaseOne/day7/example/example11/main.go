package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("%s%s%s%s", p.path, p.op, p.createTime, p.message)
}

func Open(filename string) error {
	file, err := os.Open("test.sls")
	if err != nil {
		return &PathError{
			path:       "test.sls",
			op:         "read",
			createTime: fmt.Sprintf("%v", time.Now()),
			message:    err.Error(),
		}
	}
	defer file.Close()
	return nil
}

func main() {

	//var errNotFound error = errors.New("Not found error")
	//fmt.Println(errNotFound)
	//err :=Open("c:/abc.txt")
	//v,ok :=err.(*PathError)
	//
	//if ok{
	//	fmt.Printf("111",v)
	//}
	var arr []int
	arr = make([]int, 6)
	arr = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	fmt.Println(&arr)
	arr = arr[:]
	fmt.Println(&arr)
	fmt.Println(arr)


}
