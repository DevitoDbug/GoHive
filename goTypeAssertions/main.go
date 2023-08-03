package main

import "fmt"

type Expense interface {
	cost() (amount float64)
}

type sms struct {
	body          string
	toPhoneNumber string
	isSubscribed  bool
}

func (s sms) cost() (amount float64) {
	if !s.isSubscribed {
		return float64(len(s.body)) * 10
	}
	return float64(len(s.body)) * 5
}

type email struct {
	body         string
	isSubscribed bool
	toAddress    string
}

func (s email) cost() (amount float64) {
	if !s.isSubscribed {
		return float64(len(s.body)) * .05
	}
	return float64(len(s.body)) * .01
}

func getExpenseReport(e Expense) (destination string, cost float64) {
	switch doc := e.(type) {
	case email:
		return doc.toAddress, doc.cost()
	case sms:
		return doc.toPhoneNumber, doc.cost()
	default:
		return "Error", 0
	}
}

func main() {
	emails := []email{
		{
			body:         "This is email 1",
			isSubscribed: false,
			toAddress:    "Kisumu",
		}, {
			body:         "This definitely is email 2. ",
			isSubscribed: true,
			toAddress:    "Nairobi",
		},
	}
	shortMessages := []sms{
		{
			body:          "This is the first message we have",
			toPhoneNumber: "07345",
			isSubscribed:  true,
		},
		{
			body:          "This is the second message we have",
			toPhoneNumber: "011435",
			isSubscribed:  false,
		},
	}

	fmt.Println("\nThe email report")
	for _, email := range emails {
		destination, cost := getExpenseReport(email)
		fmt.Printf("\nEmail to: %v will cost: %.2f", destination, cost)
	}
	fmt.Println("\n**********************************************************")
	fmt.Println("\nThe sms report")
	for _, message := range shortMessages {
		destination, cost := getExpenseReport(message)
		fmt.Printf("\nMessage to: %v will cost: %.2f", destination, cost)
	}
	fmt.Println("\n**********************************************************")

}
