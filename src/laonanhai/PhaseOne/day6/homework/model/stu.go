package model

import (
	"errors"
)

var ErrNotFoundBook = errors.New("err not found book")

type Student struct {
	Name  string
	Grade string
	Id    string
	Sex   string
	books []*BorrowItem
}

/**
借书记录
*/
type BorrowItem struct {
	book *Book
	num  int
}

func CreateStudent(name string, grade, id, sex string) *Student {
	return &Student{
		Name:  name,
		Grade: grade,
		Id:    id,
		Sex:   sex,
	}
}

func (s *Student) AddBook(bi *BorrowItem) {
	s.books = append(s.books, bi)
	return
}
func (s *Student) DelBook(bi *BorrowItem) (err error) {
	for i := 0; i < len(s.books); i++ {
		if s.books[i].book.Name == bi.book.Name {
			if s.books[i].num == bi.num {
				s.books = append(s.books[0:i], s.books[i+1:]...)
				return
			}
			s.books[i].num -= bi.num
		}
	}
	err = ErrNotFoundBook
	return

}
func (s *Student) GetBookList() []*BorrowItem  {
	return s.books
	
}
