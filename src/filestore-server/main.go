package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main(){
	//绑定路由
	http.HandleFunc("/file/upload",handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)
	//监听端口，绑定所有的网卡的8080端口
	err := http.ListenAndServe(":8080",nil)
	if err !=nil{
		fmt.Printf("Failed to start server ,error :%s",err.Error())
	}
}
