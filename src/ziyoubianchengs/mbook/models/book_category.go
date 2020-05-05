package models

type CategoryBook struct {
	Id int `orm:"pk;auto"json:"id"`
	BookId int
	CategoryId int
}

func (m *CategoryBook) TableName() string {
	return TNBookCategory()
}