package main

import "fmt"

type Item interface {
	GetName() string
	GetPrice() float64
}

type Product struct {
	name  string
	price float64
}

func (p Product) GetName() string {
	return p.name
}
func (p Product) GetPrice() float64 {
	return p.price
}

type Service struct {
	name  string
	price float64
}

func (s Service) GetName() string {
	return s.name
}
func (s Service) GetPrice() float64 {
	return s.price
}

func CalculateTotalPrice(i []Item) (TotalPrice float64) {
	for _, item := range i {
		TotalPrice += item.GetPrice()
	}
	return TotalPrice
}

func main() {
	//Creating dummy data

	products := []Item{
		Product{
			name:  "Milk",
			price: 400,
		},
		Product{
			name:  "Soda",
			price: 200,
		},
	}

	services := []Item{
		Service{
			name:  "Shaving",
			price: 200,
		},
		Service{
			name:  "Cleaning",
			price: 1200,
		},
	}

	//calculating the total costs
	totalProductsCost := CalculateTotalPrice(products)
	totalServiceCost := CalculateTotalPrice(services)

	fmt.Printf("The total cost for each is:\n Services: %v\n Products: %v", totalServiceCost, totalProductsCost)
}
