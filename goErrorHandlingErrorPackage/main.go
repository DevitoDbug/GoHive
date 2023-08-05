package main

import (
	"errors"
	"fmt"
	"strings"
)

func divide(dividend, divisor int) (Ans float64, err error) {
	if dividend == 0 {
		return 0, errors.New("cannot divide by zero ")
	}
	return float64(divisor / dividend), nil
}

func main() {
	repeat := "y"
	divisor := 0
	dividend := 0

	for repeat == "y" {
		fmt.Print("\nEnter divisor: ")
		_, err := fmt.Scan(&divisor)
		if err != nil {
			fmt.Print(err)
		}

		fmt.Printf("\nEnter dividend: ")
		_, err = fmt.Scan(&dividend)
		if err != nil {
			fmt.Print(err)
		}

		ans, err := divide(dividend, divisor)
		if err != nil {
			fmt.Print("\n", err)
		}
		fmt.Printf("=> %.4f", ans)

		fmt.Print("\n########################################################")
		fmt.Print("\nDo you want to continue? Enter (y/n): ")
		_, err = fmt.Scan(&repeat)
		if err != nil {
			fmt.Printf("%s\n", err)
			repeat = "n"
		}
		repeat = strings.TrimSpace(repeat)
	}

}
