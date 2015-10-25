package model

type Order struct {
	IdCustomer int
	Items []*Product
}

type Product struct {
	Id int
	Name string
	Quantity int
	Cost float64
}

type OrderReference struct {
	Id int
	IdCustomer int
	Total float64
}