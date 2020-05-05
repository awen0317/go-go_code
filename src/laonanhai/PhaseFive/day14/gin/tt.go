package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	//1.创建路由
	r :=gin.Default()
	//2.绑定路由规则
	//gin.Context,封装request和response
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK,"hello world!")
	})
	//3.监听端口
	r.Run(":8000")
}
