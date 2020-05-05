package http_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
res,err :=http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("get err",err)
		return
	}
	//提问集中夫open reade，ioutil，bufio的区别
	data,err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get data err",err)
		return
	}
	fmt.Println(string(data))

}
