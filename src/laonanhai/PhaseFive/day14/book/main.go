package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化数据
func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	// 加载页面
	r.LoadHTMLGlob("/book/template/*")
	// 查询所有图书
	r.GET("/book.list", bookListHandler)
}
func bookListHandler(c *gin.Context) {
	bookList, err := queryAllBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
	}
	c.HTML(http.StatusOK,"book_list.html",gin.H{
		"code":0,
		"data":bookList,
	})
}
