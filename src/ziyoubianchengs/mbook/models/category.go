package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id     int
	Pid    int
	Title  string
	Intro  string
	Icon   string
	Cnt    int
	Sort   int
	Status bool
}

func (m *Category) TableName() string {
	return TNCategory()
}

func (m *Category) GetCates(pid int, status int) (cates []Category, err error) {
	//创建查询条件
	qs := orm.NewOrm().QueryTable(TNCategory())

	if pid > -1 {
		qs = qs.Filter("pid", pid)
	}
	if 0 == status || 1 == status {
		qs = qs.Filter("status", status)
	}
	_, err = qs.OrderBy("-status", "sort", "title").All(&cates)
	return
}
func (m *Category) Find(cid int) (cate Category) {
	cate.Id = cid
	return cate
}
