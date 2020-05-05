package main

type Book struct {
	ID int64 `db:id`
	Title int64 `dn:title`
	Price float64 `dn:price`
}
