package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls =[]string{
	"https://www.baidu.com/",
	"https://www.google.com/",
	"https://www.taobao.com/",

}
func main()  {
	for _, v := range urls{
		//自己实现超时时间限制
		c :=http.Client{
			Transport:     &http.Transport{
				IdleConnTimeout:        2*time.Second,
			},

		}
		// resp,err :=http.Head(v)
		resp,err :=c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed,err %v",v,err)
			continue
		}
		fmt.Printf("head succ %v",resp.Status)
	}

}
