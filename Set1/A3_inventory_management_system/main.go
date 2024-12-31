package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func addProduct(products []Product, id int, name string, priceInput string, stock int) ([]Product, error) {
	price, err := strconv.ParseFloat(priceInput, 64)
	if err != nil {
		return products, fmt.Errorf("invalid price input")
	}

	product := Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
	products = append(products, product)
	return products, nil
}

func updateStock(products []Product, id int, newStock int) ([]Product, error) {
	for i := 0; i < len(products); i++ {
		if products[i].ID == id {
			if newStock < 0 {
				return products, fmt.Errorf("stock cannot be negative")
			}
			products[i].Stock = newStock
			return products, nil
		}
	}
	return products, fmt.Errorf("product not found")
}

func searchProduct(products []Product, searchTerm string) (*Product, error) {
	for _, product := range products {
		if strconv.Itoa(product.ID) == searchTerm || product.Name == searchTerm {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("product not found")
}

func displayInventory(products []Product) {
	fmt.Println("ID\tName\t\tPrice\tStock")
	fmt.Println("----------------------------------------")
	for _, product := range products {
		fmt.Printf("%d\t%s\t%.2f\t%d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func sortByPrice(products []Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
}

func sortByStock(products []Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Stock < products[j].Stock
	})
}

func main() {
	var products []Product

	products, _ = addProduct(products, 1, "Laptop", "1200.50", 10)
	products, _ = addProduct(products, 2, "Smartphone", "800.75", 20)

	products, err := updateStock(products, 1, 15)
	if err != nil {
		fmt.Println(err)
	}

	product, err := searchProduct(products, "Laptop")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Product found: %+v\n", *product)
	}

	displayInventory(products)

	sortByPrice(products)
	fmt.Println("\nSorted by Price:")
	displayInventory(products)

	sortByStock(products)
	fmt.Println("\nSorted by Stock:")
	displayInventory(products)
}
