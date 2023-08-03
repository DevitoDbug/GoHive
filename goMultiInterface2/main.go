package main

import "fmt"

type Expense interface {
	cost() float64
}
type Printer interface {
	print()
}

type Email struct {
	isSubscribed bool
	body         string
}

func (e Email) cost() (amount float64) {
	if !e.isSubscribed {
		return 0.5 * float64(len(e.body))
	}
	return 0.1 * float64(len(e.body))
}
func (e Email) print() {
	fmt.Print(e.body)
}

func CalculatePrice(e Expense, p Printer) {
	//print out the message
	//print out the price
	p.print()
	fmt.Printf("\nThe cost is: %v\n", e.cost())
	fmt.Println("====================================================")
}

func main() {
	emArr := []Email{
		{
			isSubscribed: false,
			body:         "This is email 1",
		},

		{
			isSubscribed: false,
			body:         "This is email 2",
		},
		{
			isSubscribed: false,
			body:         "This is email 3",
		},
		{
			isSubscribed: false,
			body:         "This is email 4",
		},
	}

	for _, email := range emArr {
		CalculatePrice(email, email)
	}

}
