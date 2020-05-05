package model

import (
	"errors"
)

var ErrStockNotEnough  = errors.New("stock is not enough")

type Book struct {
	Name string
	Total int
	Author string
	CreateTime string
}

func createBook(name string,total int,author string,createTime string)(b *Book)  {
	return &Book{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateTime: createTime,
	}
}
/**
是否能借书
 */
func (b *Book) canBorrow(c int) bool  {
	return b.Total >= c
}
/**
借书
 */
func (b *Book) Borrow(c int) (err error){
	if b.canBorrow(c) == false{
		err = ErrStockNotEnough
		return
	}
	b.Total -= c
	return
}
/**
还书
 */
func (b *Book) Back(c int) (err error){
	b.Total += c
	return
}