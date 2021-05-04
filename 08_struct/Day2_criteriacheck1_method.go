package main

import (
	"fmt"
	"sort"
)

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

type Products []Product

/*
type MyInt int
func (value MyInt) whoAmI() {
	fmt.Printf("I am an integer with value = %d\n", value)
}
*/

func main() {
	p1 := Product{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"}
	//Print(p1)
	p1.Print()

	products := Products{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationary"},
		{id: 107, name: "Ken", cost: 12, units: 20, category: "utencil"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "utencil"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationary"},
	}

	fmt.Println("Initial list of products")
	//Print(products)
	products.Print()

	sortProductCriteria := func(p1 *Product, p2 *Product, item string) bool {
		switch item {
		case "name":
			if p1.name > p2.name {
				p1.name, p2.name = p2.name, p1.name
			}
		case "cost":
			if p1.cost > p2.cost {
				p1.cost, p2.cost = p2.cost, p1.cost
		}
	}

	products.sortMe(sortProductCriteria)

	stationaryProductCriteria := func(product Product) bool {
		return product.category == "stationary"
	}

	anyStationaryProducts := products.Any(stationaryProductCriteria)

	fmt.Println("Any stationary products ? : ", anyStationaryProducts)

	anyGroceryProducts := products.Any(func(product Product) bool {
		return product.category == "grocery"
	})

	fmt.Println("Any grocery products ? : ", anyGroceryProducts)

	fmt.Println("All products stationary products ? :", products.All(func(product Product) bool {
		return product.category == "stationary"
	}))

	stationaryProducts := products.Filter(func(product Product) bool {
		return product.category == "stationary"
	})
	fmt.Println("Stationary Products ")
	stationaryProducts.Print()

}

func (product Product) Print() {
	fmt.Printf("Id = %d, Name = %s, cost = %v, units = %d, category = %s\n", product.id, product.name, product.cost, product.units, product.category)
}


func (products Products) Print() {
	for _, product := range products {
		product.Print()
	}
}

func (products Products) Index(product Product) int {
	for idx, productInList := range products {
		if product == productInList {
			return idx
		}
	}
	return -1
}

func (products Products) Include(product Product) bool {
	return products.Index(product) != -1
}

func (products Products) Any(criteria func(Product) bool) bool {
	for _, product := range products {
		if criteria(product) {
			return true
		}
	}
	return false
}

func (products Products) All(criteria func(Product) bool) bool {
	for _, product := range products {
		if !criteria(product) {
			return false
		}
	}
	return true
}

func (products Products) Filter(criteria func(Product) bool) Products {
	result := Products{}
	for _, product := range products {
		if criteria(product) {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) sortMe(criteria func(*Product, *Product, string) bool) bool {

	}