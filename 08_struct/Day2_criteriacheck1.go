package main

import "fmt"

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

/*
type PerishableProduct struct {
	product Product
	expiry  int
}
*/

type PerishableProduct struct {
	Product
	expiry int
}

func main() {
	/*
		var p1 Product = Product{}
		p1.id = 100
		p1.name = "Pen"
		p1.cost = 10
		p1.units = 200
		p1.category = "Stationary"
	*/
	/*
		var p1 Product = Product{100, "Pen", 10, 200, "Stationary"}
	*/
	/*
		var p1 Product = Product{name: "Pen", id: 100, units: 200, cost: 10, category: "stationary"}
	*/
	p1 := Product{name: "Pen", id: 100, units: 200, cost: 10, category: "stationary"}
	p2 := Product{name: "Pen", id: 100, units: 200, cost: 10, category: "stationary"}

	if p1 == p2 {
		fmt.Println("p1 is equal to p2")
	} else {
		fmt.Println("p1 is not equal to p2")
	}

	fmt.Printf("Type of p1 = %T\n", p1)
	fmt.Println("Product p1 = ", p1)
	applyDiscount(&p1, 10)
	fmt.Println("After applying 10% discount p1 = ", p1)

	//grapes := PerishableProduct{product: Product{101, "Grapes", 80, 50, "Fruits"}, expiry: 2}
	//fmt.Println("Perishable Product grapes = ", grapes)

	//grapes := PerishableProduct{Product: Product{101, "Grapes", 80, 50, "Fruits"}, expiry: 2}
	//grapes := PerishableProduct{Product{101, "Grapes", 80, 50, "Fruits"}, 2}
	grapes := NewPerishableProduct(101, "Grapes", 80, 50, "Fruits", 2)
	printPerishableProduct(grapes)

	products := []Product{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationary"},
		{id: 107, name: "Ken", cost: 12, units: 20, category: "utencil"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "utencil"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationary"},
	}

	stationaryProductCriteria := func(product Product) bool {
		return product.category == "stationary"
	}

	anyStationaryProducts := Any(products, stationaryProductCriteria)

	fmt.Println("Any stationary products ? : ", anyStationaryProducts)

	anyGroceryProducts := Any(products, func(product Product) bool {
		return product.category == "grocery"
	})

	fmt.Println("Any grocery products ? : ", anyGroceryProducts)

	fmt.Println("All products stationary products ? :", All(products, func(product Product) bool {
		return product.category == "stationary"
	}))

	stationaryProducts := Filter(products, func(product Product) bool {
		return product.category == "stationary"
	})
	fmt.Println("Stationary Products ")
	printProducts(stationaryProducts)

}

func printProducts(products []Product) {
	for _, product := range products {
		CustomPrint(product)
	}
}

func CustomPrint(product Product) {
	fmt.Printf("Id = %d, Name = %s, cost = %v, units = %d, category = %s\n", product.id, product.name, product.cost, product.units, product.category)
}

func Index(products []Product, product Product) int {
	for idx, productInList := range products {
		if product == productInList {
			return idx
		}
	}
	return -1
}

func Include(products []Product, product Product) bool {
	return Index(products, product) != -1
}

func Any(products []Product, criteria func(Product) bool) bool {
	for _, product := range products {
		if criteria(product) {
			return true
		}
	}
	return false
}

func All(products []Product, criteria func(Product) bool) bool {
	for _, product := range products {
		if !criteria(product) {
			return false
		}
	}
	return true
}

func Filter(products []Product, criteria func(Product) bool) []Product {
	result := []Product{}
	for _, product := range products {
		if criteria(product) {
			result = append(result, product)
		}
	}
	return result
}

/*
func AnyProductWithUnitsGt50(products []Product) bool {
	for _, product := range products {
		if product.units > 50 {
			return true
		}
	}
	return false
}
func AnyStationaryProducts(products []Product) bool {
	for _, product := range products {
		if product.category == "stationary" {
			return true
		}
	}
	return false
}
*/

/*
func printPerishableProduct(pp PerishableProduct) {
	fmt.Printf("Perishable Product => id=%d, name=%s, cost=%v, units=%d, category=%s, expiry=%d\n", pp.product.id, pp.product.name, pp.product.cost, pp.product.units, pp.product.category, pp.expiry)
}
*/

/*
func printPerishableProduct(pp PerishableProduct) {
	fmt.Printf("Perishable Product => id=%d, name=%s, cost=%v, units=%d, category=%s, expiry=%d\n", pp.Product.id, pp.Product.name, pp.Product.cost, pp.Product.units, pp.Product.category, pp.expiry)
}
*/

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry int) PerishableProduct {
	return PerishableProduct{Product{id, name, cost, units, category}, expiry}
}

func printPerishableProduct(pp PerishableProduct) {
	fmt.Printf("Perishable Product => id=%d, name=%s, cost=%v, units=%d, category=%s, expiry=%d\n", pp.id, pp.name, pp.cost, pp.units, pp.category, pp.expiry)
}

func applyDiscount(productPtr *Product, discount int) /*  */ {
	productPtr.cost = productPtr.cost * ((100 - float32(discount)) / 100)
}
