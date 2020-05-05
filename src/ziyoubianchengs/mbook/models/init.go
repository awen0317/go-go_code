package models

import (
	"github.com/astaxie/beego/orm"
)

func init()  {
	orm.RegisterModel(
		new(Category),
		new(Book),
		new(CategoryBook),
		)
}

func TNCategory() string{
	return "md_category"
}
func TNBook() string{
	return "md_books"
}
func TNBookCategory() string {
	return "md_book_category"
}