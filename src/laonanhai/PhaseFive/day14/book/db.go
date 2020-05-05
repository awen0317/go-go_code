package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func initDB() (err error) {
	addr := "root:123@tcp(127.0.0.1:3600)/test"
	db,err = sqlx.Connect("mysql",addr)
	if err != nil {
		return
	}
	//最大连接数
	db.SetMaxOpenConns(100)
	//最大连最大空闲
	db.SetMaxIdleConns(16)
	return
}
func queryAllBook()(bookList []*Book,err error)  {
	sqlStr :="select id,title,price from book"
	err  =db.Select(&bookList,sqlStr)
	if err != nil {
		fmt.Println("查询失败")
		return
	}
	return
}
func insertBook(title string,price float64)(err error)  {

	sqlStr := "insert into book(title,price) value (?,?)"
	_,err =db.Exec(sqlStr,title,price)
	if err != nil {
		fmt.Println("插入失败")
		return
	}
	return
}
func deleteBook(id int64)(err error) {
	sqlStr := "delete from book where id = ?"
	_,err =db.Exec(sqlStr,id)

	if err != nil {
		fmt.Println("删除失败")
		return
	}
	return


}