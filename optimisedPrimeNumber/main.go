package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	var isPrime bool
	var repeate bool = true

	for repeate {
		fmt.Print("Enter a number to evaluate: ")
		fmt.Scan(&number)

		if isPrime = evaluatePrimeNumber(number); isPrime {
			fmt.Printf("The number %d is prime\n", number)
		} else {
			fmt.Printf("The number %d is not prime\n", number)
		}

		fmt.Print("Do you want to continue? (y/n): ")
		var input string
		fmt.Scan(&input)
		if input == "y" {
			repeate = true
		} else if input == "n" {
			repeate = false
		} else {
			fmt.Println("Invalid input")
			repeate = false
		}

	}

}

func evaluatePrimeNumber(number int) bool {
	//Checking if it is less than 2
	if number <= 1 {
		return false
	}

	//Checking if it is even
	if number%2 == 0 && number != 2 {
		return false
	}

	//Checking if it is odd number
	if number%3 == 0 && number != 3 {
		return false
	}

	sqrt := int(math.Sqrt(float64(number)))

	//incrimenting by 6 to skip even and odd numbers
	for i := 5; i <= sqrt; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}
