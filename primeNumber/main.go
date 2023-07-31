package main

import (
	"fmt"
	"math"
)

func main() {
	num := 0
	isPrime := true
	displayText := map[bool]string{
		true:  "Is a prime number",
		false: "Is not a prime number",
	}
	repeat := true
	for repeat {
		fmt.Print("Enter a number: ")
		fmt.Scan(&num)

		//Check if it is a prime number
		rootNum := math.Sqrt(float64(num))
		for i := 2; i <= int(rootNum); i++ {
			if num%i == 0 && num != 3 {
				isPrime = false
				break
			}
		}

		fmt.Printf("%v - %s ", num, displayText[isPrime])

		fmt.Print("\n\n To check another number enter 0 else 1:  ")
		fmt.Scan(&repeat)

	}
}
