package main

import (
	"fmt"
)

// single method interface
type PriceCalculator interface {
	CalculateTotalCost() (cost float32)
}

type Product struct {
	Id    int32
	Name  string
	Price float32
}

type Order struct {
	Id          int32
	ProductList []Product
}

func (o *Order) AddProduct(p Product) {
	o.ProductList = append(o.ProductList, p)
}

func (o *Order) CalculateTotalCost() (cost float32) {
	for _, p := range o.ProductList {
		cost += p.Price
	}
	return
}

func main() {
	o := Order{Id: 1}
	o.AddProduct(Product{Id: 1, Name: "milk", Price: 1.1})
	// method value
	addProduct := o.AddProduct
	addProduct(Product{Id: 2, Name: "bread", Price: 2.11})

	// method expression
	calculateTotalCost := (*Order).CalculateTotalCost
	// the first parameter is the receiver of the method
	cost := calculateTotalCost(&o)
	fmt.Printf(" Total Cost : $%f\r\n", cost)

	addProduct(Product{Id: 3, Name: "bread", Price: 2.11})
	// method expression using interface
	calculatePrice := PriceCalculator.CalculateTotalCost
	// the first parameter is always the receiver of the method
	cost = calculatePrice(&o)
	fmt.Printf(" Total Cost : $%f\r\n", cost)

	o2 := Order{Id: 2}
	o.AddProduct(Product{Id: 1, Name: "milk", Price: 1.1})
	cost = calculatePrice(&o2)
	fmt.Printf(" Total Cost : $%f\r\n", cost)
}
